package main

import (
	"encoding/csv"
	"fmt"
	"github.com/GeertJohan/go.ask"
	"github.com/conformal/btcchain"
	"os"
	"strconv"
	"strings"
)

// export handles exporting the chain to a file with filename
// export assumes control of communication with user
func export(sim *Simulator, filename string) {
	if !strings.HasSuffix(filename, ".csv") {
		filename = filename + ".csv"
	}
	// see if file exists and if so: ask user if we should overwrite
	_, err := os.Stat(filename)
	if err == nil {
		overwrite := ask.MustAskf("file `%s` exists, overwrite?", filename)
		if !overwrite {
			fmt.Println("Did not export chain to file!")
			return
		}
	} else if !os.IsNotExist(err) {
		fmt.Printf("Did not export chain to file! Error: %v\n", err)
		return
	}

	// setup csv file writer
	file, err := os.Create(filename)
	if err != nil {
		fmt.Print("Did not export chain to file! Error creating file: %v\n", err)
		return
	}
	csvWriter := csv.NewWriter(file)

	// setup data
	data := make([][]string, 0, int(sim.chain.height())+4)
	// line 1: mining simulation used
	data = append(data, []string{sim.miningSim.Name()})
	// line 2: difficulty algorithm used
	data = append(data, []string{sim.algo.Name()})
	// line 3 ?
	data = append(data, []string{""})
	// line 4: column headers
	data = append(data, []string{"height", "diff compact", "diff human", "seconds since prev", "duration since prev", "timestamp"})

	// add chain to [][]string
	var prevTime = sim.chain.getBlock(0).time
	for i := int64(0); i < sim.chain.height(); i++ {
		block := sim.chain.getBlock(i)
		timeDiff := block.time.Sub(prevTime)
		data = append(data, []string{
			strconv.Itoa(int(i)),                                 // block ID
			fmt.Sprintf("%x", btcchain.BigToCompact(block.diff)), // difficulty compacted bits??
			fmt.Sprintf("%.10f", DiffToHumanFloat64(block.diff)), // difficulty (human representation e.g. `712.03`)
			fmt.Sprintf("%d", int(timeDiff.Seconds())),           // seconds since previous block
			timeDiff.String(),                                    // time since previous block (human representation e.g. `1m 30s`)
			block.time.Format("2006-01-02 15:04:05"),             // full timestamp
		})
		prevTime = block.time
	}

	// encode data directly to file
	err = csvWriter.WriteAll(data)
	if err != nil {
		fmt.Printf("Error while writing chain to file: %v\n", err)
		return
	}
	err = csvWriter.Error()
	if err != nil {
		fmt.Printf("Error while writing chain to file: %v\n", err)
		return
	}
}
