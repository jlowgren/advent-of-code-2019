package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	chunkSize      = 4
	opcodeAdd      = 1
	opcodeMultiply = 2
	opcodeExit     = 99
)

func main() {
	numbers := readFile("input.txt")

	numbers[1] = 12
	numbers[2] = 2

	i := 0
	for {
		if numbers[i] == opcodeExit {
			break
		}

		opcode := numbers[i]
		pos1 := numbers[numbers[i+1]]
		pos2 := numbers[numbers[i+2]]

		if opcode == opcodeAdd {
			numbers[numbers[i+3]] = pos1 + pos2
		} else if opcode == opcodeMultiply {
			numbers[numbers[i+3]] = pos1 * pos2
		} else {
			fmt.Println("Unknown opcode")
			break
		}

		i += chunkSize
	}

	fmt.Println(numbers[0])
}

func readFile(fileName string) []int {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var numbers []int
	for _, value := range strings.Split(string(input), ",") {
		number, err := strconv.Atoi(strings.TrimRight(value, "\n"))
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}
