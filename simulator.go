package main

import (
	"math/big"
	"time"
)

// MiningSimulation is implemented by several simulations such as 'simple wave' and 'multipool'.
type MiningSimulation interface {
	Name() string
	Setup(*chain, DiffAlgo)

	// Hashrate returns the hashrate that is now in the network
	Hashrate() *big.Int
}

type Simulator struct {
	chain     *chain
	miningSim MiningSimulation
	algo      DiffAlgo

	hashrate *big.Int
}

func (sim *Simulator) SimulateBlocks(n int) {

	// 30354 KH per 1 diff at 150 seconds blocktime

	// 1 diff = 30354 * 150 = 4553100

	for i := 0; i < n; i++ {
		// update network hashrate
		sim.hashrate = sim.miningSim.Hashrate()
		// calculate diff for new block
		diff := sim.algo.Calculate(sim.chain)

		// calculate time the new block takes with new hashrate and new diff
		perfectHashrate := DiffToHashratePerfect(diff)
		verbosef("perfect hashrate: %s\n", perfectHashrate.String())
		// duration := time.Duration(float64(targetSpacing) * (diff / (sim.hashrate / 4553100.0))) //++ TODO: use new formula
		durationSeconds := (&big.Int{})
		durationSeconds.Mul(targetSpacingSeconds, perfectHashrate)
		durationSeconds.Div(durationSeconds, sim.hashrate)
		duration := time.Duration(durationSeconds.Uint64()) * time.Second
		verbosef("block duration: %s\n", duration.String())
		t := sim.chain.lastBlock().time.Add(duration)

		// add new block to chain
		b := &block{
			time: t,
			diff: diff,
		}
		sim.chain.addBlock(b)
	}
}

func (sim *Simulator) GetLastDiff() *big.Int {
	return sim.chain.lastBlock().diff
}
func (sim *Simulator) GetLastHashrate() *big.Int {
	return sim.hashrate
}
