
## diffsim

diffsim is a program that simulates a bitcoin network. Its main purpose is to test existing difficulty algorithms and assist in developing a new algorithm.

### Install

First install go: [https://golang.org/doc/install](https://golang.org/doc/install)

Use the go tool to get and install the diffsim sources: `go get github.com/GeertJohan/diffsim`

### Usage

Simply run the application in a terminal. Use the `-v` flag to enable verbose logging.

```
diffsim -v
```

The application will ask you which algorithm and simulation to run.
Right now only one simulation and one algorithm is available, but more will become available soon.

### Contribute
I accept pull requests containing new simulations or difficulty re-adjustment algorithms.

If you wish to change something, please open an issue first to discuss the change.

### Example usage

```
$ diffsim
Which algorithm should be tested? 
1) DGW3
> 1

Please choose a simulation to run:
1) Simple waves
> 1

Using Dark Gravity Wave 3 algo for Simple Waves simulation
Starting simulation, enter 'help' for help on commands.
> 1
Latest block diff: 110427836236357352041769395878404723568785424659630784333489133269811200, hashrate used for last block was: 7723
> 2
Latest block diff: 110427836236357352041769395878404723568785424659630784333489133269811200, hashrate used for last block was: 7568
> 100
Latest block diff: 110427836236357352041769395878404723568785424659630784333489133269811200, hashrate used for last block was: 6472
> 1000
Latest block diff: 2777861570415274433432389690966477056895500698541819974748315312109962, hashrate used for last block was: 252086
> 100
Latest block diff: 1801263199947073992204610835070877071294746243320012125156996004921323, hashrate used for last block was: 475108
> print 5
block 1199
	timestamp:                 2014-09-03 03:40:00
	seconds since prev block:  0
	difficulty:                1781060482977257282231265323979430191767359497440689752263342294388406
block 1200
	timestamp:                 2014-09-03 03:42:35
	seconds since prev block:  2m35s
	difficulty:                1785333944578041833317606548601373415166629870710145349216233202246177
block 1201
	timestamp:                 2014-09-03 03:45:03
	seconds since prev block:  2m28s
	difficulty:                1795507688585126191766410073401290330535419712823797789548443340845402
block 1202
	timestamp:                 2014-09-03 03:47:18
	seconds since prev block:  2m15s
	difficulty:                1803659974196706935994301852216352408780612255865393239985184514931351
block 1203
	timestamp:                 2014-09-03 03:49:25
	seconds since prev block:  2m7s
	difficulty:                1805471702557075653034577726268057997067030482506000202509322211730703
block 1204
	timestamp:                 2014-09-03 03:51:40
	seconds since prev block:  2m15s
	difficulty:                1801263199947073992204610835070877071294746243320012125156996004921323
```


#### Export

What is in the exported csv file;
 - block ID
 - difficulty compacted bits??
 - difficulty human representation (e.g. 712.03)
 - seconds since previous block
 - time since previous block (human representation e.g. `1m 30s`)
 - full timestamp


### License
This project is licensed under a Simplified BSD license. Please read the [LICENSE file][license].
