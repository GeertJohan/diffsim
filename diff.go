package main

import (
	"math"
	"math/big"
)

var (
	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)
	bigK   = big.NewInt(1000)

	// oneLsh256 is 1 shifted left 256 bits.  It is defined here to avoid
	// the overhead of creating it multiple times.
	oneLsh256 = new(big.Int).Lsh(bigOne, 256)
)

// DiffToHumanFloat64 returns te human diff for a *big.Int bigmask
func DiffToHumanFloat64(diff *big.Int) float64 {
	var modifier = 1. / float64(math.Pow(2, 12)) / 1000

	diffNum := (&big.Int{}).Div((&big.Int{}).Mul(powDiffLimit, bigK), (&big.Int{}).Mul(diff, bigK))

	return float64(diffNum.Uint64()) * modifier
}

// DiffToHashratePerfect calculates the required hashrate to mine one block in a perfect-world situation (150 seconds)
func DiffToHashratePerfect(diff *big.Int) *big.Int {
	hashesPerBlock := (&big.Int{}).Div(oneLsh256, diff)
	hashrate := (&big.Int{}).Div(hashesPerBlock, targetSpacingSeconds)
	return hashrate
}
