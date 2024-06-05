package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// reads the file at argument[0] line by line and converts them to a slice of float64
func Readfile(str string) []float64 {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []float64

	data := 0.0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}
		ReplaceFiles := strings.ReplaceAll(input, ",", "")
		TrimFile := strings.TrimSpace(ReplaceFiles)
		data, err = strconv.ParseFloat(TrimFile, 64)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, data)

	}

	if len(output) == 0 {
		log.Fatal("Empty file")
	}

	return output

}
