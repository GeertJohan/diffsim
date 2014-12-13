package main

import (
	"math/big"
	"math/rand"
)

type SimpleWaveSimulator struct {
	chain               *chain
	algo                DiffAlgo
	hashrate            *big.Int // hashrate in hash/s
	maxChangePercentage int32
}

func NewSimpleWaveSimulator() *SimpleWaveSimulator {
	return &SimpleWaveSimulator{
		hashrate:            big.NewInt(6990), // initial hashrate of 6990 hash/sec which should solve 1 block in 150 seconds...
		maxChangePercentage: 10,               // maximum hashrate change in percent between blocks
	}
}

func (sws *SimpleWaveSimulator) Name() string {
	return "Simple Waves"
}

func (sws *SimpleWaveSimulator) Setup(c *chain, a DiffAlgo) {
	sws.chain = c
	sws.algo = a
}

// Hashrate returns the current hashrate
// SimpleWaveSimulator disregards the hashrate
func (sws *SimpleWaveSimulator) Hashrate() *big.Int {
	// get a random between 0 and 2*maxChangePercentage
	// NOTE: random will alway follow the same pattern in Go because math/rand is not seeded.
	// 		This means that although the numbers returned by rand.Int21n are pseudo-random, they are predictable and the same for each run.
	rnd := rand.Int31n(sws.maxChangePercentage*2 + 1)

	// calculate modifier based on random percentageChange
	modifier := 1.005 + (float64(rnd-sws.maxChangePercentage) / 100.0) // let hashrate slowly increase
	sws.hashrate.Div(sws.hashrate.Mul(sws.hashrate, big.NewInt(int64(modifier*100000.0))), big.NewInt(100000))

	// don't let hashrate come too low.
	if sws.hashrate.Cmp(big.NewInt(100)) < 0 {
		sws.hashrate.Set(big.NewInt(100))
	}
	verbosef("random hashrate change: %.2f, new hashrate: %s\n", modifier, hashrateToString(float64(sws.hashrate.Uint64())))
	return sws.hashrate
}
