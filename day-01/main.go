package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	numbers := readFile("input.txt")

	total := 0
	for _, mass := range numbers {
		fuel := (mass / 3) - 2
		total += fuel
	}

	fmt.Println(total)
}

func readFile(fileName string) []int {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var numbers []int
	for _, value := range strings.Fields(string(input)) {
		number, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}
