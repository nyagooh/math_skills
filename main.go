package main

import (
	"fmt"
	"log"
	"math"
	m "mathskills/functions"
	"os"
	"strings"
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
	input := m.Readfile(args1)
	mean := math.Round(m.CalculateMean(input))
	median := math.Round(m.CalculateMedian(input))
	variance := int(math.Round(m.CalculateVariance(input)))
	deviation := math.Round(math.Sqrt(m.CalculateVariance(input)))

	fmt.Printf("Average: %v\nMedian: %v\nVariance:%v\nStandard Deviation:%v", mean, median, variance, deviation)
	fmt.Println()

}