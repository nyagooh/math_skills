package main

import (
	"fmt"
	"log"
	m "mathskills/functions"
	"os"
	"strings"
	"math"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("not enough arguments passed")
	}
	args1 := args[0]
	isSuffix := strings.HasSuffix(args1, "txt")
	if !isSuffix {
		log.Fatal("wrong extension at args1")
	}
	input := m.Readfile()
	mean := math.Round(m.CalculateMean(input))
	median := math.Round(m.CalculateMedian(input))
	deviation := math.Round(m.CalculateStandardSDeviation(input))
	variance := math.Round(m.CalculateVariance(input))
	fmt.Printf("mean%vmedian%vdeviation%vvariance%v",mean,median,deviation,variance)
	fmt.Println()

}
