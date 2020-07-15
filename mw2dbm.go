package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	// array of float args
	var floatArgs []float64

	// sum of all mw input values
	var mwSum float64
	var dbmSum float64

	// grab all CLI args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("zero!")
		os.Exit(1)
	}

	// convert all CLI args to floats
	for _, val := range args {
		tmp, _ := strconv.ParseFloat(val, 8)
		floatArgs = append(floatArgs, tmp)
		mwSum += tmp
	}

	// if there's more than one arg, print a table
	for _, val := range floatArgs {
		dbm := mw2dbm(val)
		fmt.Printf("mw %6.4f db %6.4f\n", val, dbm)
	}

	if len(floatArgs) > 1 {
		dbmSum = mw2dbm(mwSum)
		fmt.Println("mw sum", mwSum)
		fmt.Println("dbm sum", dbmSum)
	}

}

func mw2dbm(mw float64) float64 {
	// receive mw
	// return int
	return math.Log10(mw) * 10

}
