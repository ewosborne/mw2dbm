package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	// grab all CLI args
	args := os.Args[1:]

	// error checks
	if len(args) == 0 {
		fmt.Println("zero args given! that's dumb!")
		os.Exit(1)
	}

	// convert all CLI args to floats
	mwArgs := arrayToFloat(args)

	// convert everything in mwArgs to dBm
	dbmArgs := mwArrayToDbm(mwArgs)

	// print columns of mw and dbm
	printSummaryTable(mwArgs, dbmArgs)

}

func mwArrayToDbm(items []float64) []float64 {
	var retVal []float64
	for _, val := range items {
		retVal = append(retVal, mw2dbm(val))
	}
	return retVal
}

func mw2dbm(mw float64) float64 {
	// receive mw
	// return dbm
	return math.Log10(mw) * 10

}

func sumFloat64(items []float64) float64 {
	var arraySum float64
	for _, val := range items {
		arraySum += val
	}
	return arraySum
}

func arrayToFloat(src []string) []float64 {
	var floatArgs []float64
	for _, val := range src {
		tmp, _ := strconv.ParseFloat(val, 8)
		floatArgs = append(floatArgs, tmp)
	}
	return floatArgs
}

func printSummaryTable(mwArgs []float64, dbmArgs []float64) {
	fmt.Println("        mw | dbm ")
	for idx, val := range mwArgs {
		fmt.Printf("    %6.3f | %6.3f\n", val, dbmArgs[idx])
	}

	// if there's more than one argument, also print the column sums
	if len(mwArgs) > 1 {
		// sum up all mw
		mwSum := sumFloat64(mwArgs)
		dbmSum := mw2dbm(sumFloat64(mwArgs))
		fmt.Printf("SUM:%6.3f | %6.3f\n", mwSum, dbmSum)

	}
}
