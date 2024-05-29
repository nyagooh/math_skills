package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)
// reads the file at argument[0] line by line
func Readfile(str string) []float64 {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []float64

	data := 0.0
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		input := scanner.Text()
		data, err = strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, data)
		count++
	}

	if count == 0 {
		log.Fatal("Empty file")
	}

	return output

}
