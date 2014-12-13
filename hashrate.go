package main

import (
	"fmt"
	"math"
)

func hashrateToString(hashrate float64) string {
	switch true {
	case hashrate > math.Pow(1000, 4):
		return fmt.Sprintf("%.3f TH/s", hashrate/float64(math.Pow(1000, 4)))
	case hashrate > math.Pow(1000, 3):
		return fmt.Sprintf("%.3f GH/s", hashrate/float64(math.Pow(1000, 3)))
	case hashrate > math.Pow(1000, 2):
		return fmt.Sprintf("%.3f MH/s", hashrate/float64(math.Pow(1000, 2)))
	case hashrate > math.Pow(1000, 1):
		return fmt.Sprintf("%.3f KH/s", hashrate/float64(math.Pow(1000, 1)))
	default:
		return fmt.Sprintf("%.3f hash/s", hashrate)
	}
}
