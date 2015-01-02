package main

import (
	"github.com/conformal/btcchain"
	"math/big"
	"time"
)

// DarkGravityWave3 implements the DiffAlgo interface.
// It performs the DarkGravityWave algorithm version 3 as initially implemented in C++ by Evan Duffield - evan@darkcoin.io
type DarkGravityWave3 struct {
}

func NewDarkGravityWave3() *DarkGravityWave3 {
	return &DarkGravityWave3{}
}

func (dgw *DarkGravityWave3) Name() string {
	return "Dark Gravity Wave 3"
}

// Calculate calculates the new DGW3 difficulaty based on the current chain
func (dgw *DarkGravityWave3) Calculate(chain *chain) *big.Int {
	var actualTimespan time.Duration
	var targetTimespan time.Duration
	pastBlocksMin := int64(24)
	pastBlocksMax := int64(24)
	countBlocks := int64(0)

	curHeight := chain.height()
	if curHeight < pastBlocksMin {
		verbosef("DGW3 requires %d blocks beofre calculations, returning diff limit (%.10f)\n", pastBlocksMin, DiffToHumanFloat64(powDiffLimit))
		return (&big.Int{}).Set(powDiffLimit)
	}

	// loop over the past n blocks, where n == PastBlocksMax
	var pastDifficultyAverage *big.Int
	var prevBlockTime time.Time
	for i := int64(0); i < pastBlocksMax; i++ {
		countBlocks++

		b := chain.getBlock(curHeight - i)

		// Calculate average difficulty based on the blocks we iterate over in this for loop
		if countBlocks == 1 {
			pastDifficultyAverage = (&big.Int{}).Set(b.diff)
		} else {
			// pastDifficultyAverage = ((pastDifficultyAverage * countBlocks) + (b.diff)) / (countBlocks + 1) // WARN: this does not calculate a real average, but this is how DGW3 does it...
			// pastDifficultyAverage = ((pastDifficultyAverage * (countBlocks - 1)) + b.diff) / countBlocks // is real avg, bug not in dgw3 algo
			pastDifficultyAverage.Mul(pastDifficultyAverage, big.NewInt(countBlocks))
			pastDifficultyAverage.Add(pastDifficultyAverage, b.diff)
			pastDifficultyAverage.Div(pastDifficultyAverage, big.NewInt(countBlocks+1))
		}

		// If this is the second iteration (LastBlockTime was set)
		if i > 0 {
			// Increment the actual timespan
			actualTimespan += prevBlockTime.Sub(b.time)
		}
		// Set LasBlockTime to the block time for the block in current iteration
		prevBlockTime = b.time
	}

	newDifficulty := (&big.Int{}).Set(pastDifficultyAverage)

	// nTargetTimespan is the time that the CountBlocks should have taken to be generated.
	targetTimespan = time.Duration(countBlocks) * targetSpacing

	verbosef("targetTimespan: %.0f seconds\n", targetTimespan.Seconds())
	verbosef("actualTimespan: %.0f seconds\n", actualTimespan.Seconds())

	// Limit the re-adjustment to 3x or 0.33x
	// We don't want to increase/decrease diff too much.
	if actualTimespan < targetTimespan/3 {
		actualTimespan = targetTimespan / 3
	}
	if actualTimespan > targetTimespan*3 {
		actualTimespan = targetTimespan * 3
	}

	verbosef("targetTimespan limited: %.0f\n", targetTimespan.Seconds())
	verbosef("actualTimespan limited: %.0f\n", actualTimespan.Seconds())

	// Calculate the new difficulty based on actual and target timespan.
	newDifficulty.Mul(newDifficulty, big.NewInt(int64(actualTimespan.Seconds())))
	newDifficulty.Div(newDifficulty, big.NewInt(int64(targetTimespan.Seconds())))

	verbosef("newDifficulty without pow limit:  %x %s\n", btcchain.BigToCompact(newDifficulty), newDifficulty.String())

	// If calculated difficulty is greater than the minimal diff, set the new difficulty to be the minimal diff.
	if newDifficulty.Cmp(powDiffLimit) > 0 {
		verboseln("DGW3: setting new diff to powDiffLimit")
		newDifficulty.Set(powDiffLimit)
	}

	// Some logging.
	// TODO: only display these log messages for a certain debug option.
	verbosef("Before: %x %s\n", btcchain.BigToCompact(chain.lastBlock().diff), chain.lastBlock().diff.String())
	verbosef("After:  %x %s\n", btcchain.BigToCompact(newDifficulty), newDifficulty.String())

	// Return the new diff.
	return newDifficulty
}
