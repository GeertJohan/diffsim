package main

import (
	"math/big"
)

// DiffAlgo is implemented by DGW3 and GDR
type DiffAlgo interface {
	// Name must return the name for the interface
	Name() string

	// Calculate gives the diff algo the chain on which to calculate the new difficulty
	// The returned difficulty (*bit.Int) indicates the amount of work that must be done before a valid block hash is hit.
	Calculate(c *chain) *big.Int
}
