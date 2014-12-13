package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/GeertJohan/go.linenoise"
	"github.com/foize/go.sgr"
)

func main() {
	initFlags()
	verboseln("flags initialized")

	sim := &Simulator{}

	fmt.Println(`Which algorithm should be tested? 
1) DGW3`)
	// 2) GDR`)
	for {
		ans, err := linenoise.Line("> ")
		if err != nil {
			if err == linenoise.KillSignalError {
				fmt.Println("quiting..")
				os.Exit(0)
			}
			fmt.Printf("error reading line: %v", err)
			os.Exit(1)
		}
		switch ans {
		case "1", "DGW3":
			sim.algo = NewDarkGravityWave3()
		// case "2", "GDR":
		// 	sim.algo = NewGuldenDifficultyReadjustment()
		default:
			fmt.Printf("Unkown algorithm '%s', please try again\n", ans)
			continue
		}
		break
	}
	fmt.Println("")

	fmt.Println(`Please choose a simulation to run:
1) Simple waves`)
	for {
		ans, err := linenoise.Line("> ")
		if err != nil {
			if err == linenoise.KillSignalError {
				fmt.Println("quiting..")
				os.Exit(0)
			}
			fmt.Printf("error reading line: %v", err)
			os.Exit(1)
		}
		switch ans {
		case "1":
			sim.miningSim = NewSimpleWaveSimulator()
		default:
			fmt.Printf("Unkonwn simulator '%s', please try again\n", ans)
			continue
		}
		break
	}
	fmt.Println("")

	fmt.Printf("Using %s algo for %s simulation\n", sim.algo.Name(), sim.miningSim.Name())

	// setup chain
	sim.chain = newChain()
	t, err := time.Parse("2006-01-02 15:04:05", "2014-09-01 00:00:00")
	if err != nil {
		fmt.Printf("error parsing time: %v\n", err)
		os.Exit(1)
	}
	// add perfect history
	for sim.chain.height() < 1 {
		b := &block{
			time: t,
			diff: big.NewInt(1), // assuming initial diff of 1
		}
		sim.chain.addBlock(b)
		t = t.Add(targetSpacing)
	}

	fmt.Println("Starting simulation, enter 'help' for help on commands.")
	for {
		line, err := linenoise.Line("> ")
		if err != nil {
			if err == linenoise.KillSignalError {
				fmt.Println("quiting..")
				os.Exit(0)
			}
			fmt.Printf("error reading line: %v", err)
			os.Exit(1)
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]
		switch cmd {
		case "help":
			fmt.Println(`
help       Print this message
quit       Stop the simulation
last <n>   Print last n blocks from chain
1          Simulate a single block
<n>        Simulate n blocks`)
		case "quit":
			fmt.Println("Thanks for using diffsim")
			os.Exit(0)
		case "last":
			n, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				fmt.Printf("Invalid input '%s'.\n", line)
				continue
			}
			if int64(n) > sim.chain.height() {
				fmt.Printf("Cannot print last %d blocks, chain is only %d high\n", n, sim.chain.height())
				continue
			}
			var prevBlock *block
			for i := sim.chain.height() - int64(n); i <= sim.chain.height(); i++ {
				block := sim.chain.getBlock(i)
				var timediff time.Duration
				if prevBlock != nil {
					timediff = block.time.Sub(prevBlock.time)
				}
				fmt.Printf(`block %d
	timestamp:                 %s
	seconds since prev block:  %s
	difficulty:                %s
	difficulty representation: %.03f
`, i, block.time.Format(`2006-01-02 15:04:05`), timediff.String(), block.diff.String(), DiffToHumanFloat64(block.diff))
				prevBlock = block
			}
		default:
			n, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				fmt.Printf("Invalid input '%s'.\n", line)
				continue
			}
			verbosef("Adding %d blocks\n", n)
			sim.SimulateBlocks(int(n))
			fmt.Printf("Latest block diff: %s, hashrate used for last block was: %s\n", sim.GetLastDiff().String(), sim.GetLastHashrate().String())
		}
	}
}

func verboseln(line string) {
	if flags.Verbose {
		fmt.Println(sgr.FgBlue + line + sgr.Reset)
	}
}

func verbosef(format string, d ...interface{}) {
	if flags.Verbose {
		fmt.Printf(sgr.FgBlue+format+sgr.Reset, d...)
	}
}
