package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Readfile() []float64 {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var input string
	var output []float64
	data := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
		data, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(err)
		}
		output = append(output, data)
	}

	defer file.Close()
	return output

}
