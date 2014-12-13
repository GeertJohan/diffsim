package main

import (
	"github.com/conformal/btcchain"
	"math/big"
	"time"
)

// This file contains coin specifications (main chain).

const targetSpacing = 150 * time.Second // Guldencoin: 2.5 minutes between block

var targetSpacingSeconds = (&big.Int{}).SetInt64(int64(targetSpacing.Seconds()))

var powDiffLimit = btcchain.CompactToBig(0x1e0fffff)
