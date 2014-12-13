package main

import (
	"math/big"
	"time"
)

type chain struct {
	ladder []*block
}

func newChain() *chain {
	return &chain{
		ladder: make([]*block, 0),
	}
}

func (c *chain) addBlock(b *block) {
	b.height = int64(len(c.ladder))
	c.ladder = append(c.ladder, b)
}

func (c *chain) getBlock(height int64) *block {
	return c.ladder[height]
}

func (c *chain) lastBlock() *block {
	return c.ladder[len(c.ladder)-1]
}

func (c *chain) height() int64 {
	return int64(len(c.ladder)) - 1
}

type block struct {
	height int64
	time   time.Time
	diff   *big.Int
}
