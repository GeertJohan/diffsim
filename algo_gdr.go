package main

import (
	"math/big"
)

// GuldenDifficultyReadjustment implements DiffAlgo
// GuldenDifficultyReadjustment is a work-in-progress difficulty readjustment algorithm created by Geert-Johan Riemer and the Guldencoin community (TODO: contributors list)
type GuldenDifficultyReadjustment struct {
	chain *chain
}

// NewGuldenDifficultyReadjustment returns a new GuldenDifficultyReadjustment instance
func NewGuldenDifficultyReadjustment() *GuldenDifficultyReadjustment {
	return &GuldenDifficultyReadjustment{}
}

func (gdr *GuldenDifficultyReadjustment) Name() string {
	return "Gulden Difficulty Re-adjustment"
}

// Calculate calculates the new GDR difficulaty based on the current chain
func (gdr *GuldenDifficultyReadjustment) Calculate(chain *chain) *big.Int {
	return (&big.Int{}).Set(powDiffLimit)
	// // check if algo can start to calculate
	// pastBlocksMin := int64(24)
	// curHeight := gdr.chain.height()
	// if curHeight < pastBlocksMin {
	// 	return powDiffLimit
	// }

	// // ## TODO
	// // ## input and calculated data

	// lastBlock := gdr.chain.lastBlock()

	// var avg100Blocks = 100.0 // should be calculated

	// // ## TODO
	// // ## actual algorithm

	// var (
	// 	firstLowerBoundary  = lastBlock.diff // * 0.5 // calculation doesnt work on *big.Int
	// 	firstUpperBoundary  = lastBlock.diff // * 1.2 // calculation doesnt work on *big.Int
	// 	secondLowerBoundary = avg100Blocks   // * 3.0 // calculation doesnt work on *big.Int
	// 	secondUpperBoundary = avg100Blocks   // * 0.33 // calculation doesnt work on *big.Int
	// )

	// var factor = 1.2

	// // calculage new diff based on last diff
	// newDiff := lastBlock.diff * factor

	// // check boundary's
	// if newDiff < firstLowerBoundary {
	// 	return firstLowerBoundary
	// }
	// if newDiff > firstUpperBoundary {
	// 	return firstUpperBoundary
	// }
	// if newDiff < secondLowerBoundary {
	// 	return secondLowerBoundary
	// }
	// if newDiff > secondUpperBoundary {
	// 	return secondUpperBoundary
	// }

	// return newDiff
}
