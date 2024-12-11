package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func stringToIntArray(s string) []int {
	var intArr []int
	strings := strings.Split(s, " ")
	for _, v := range strings {
		num, _ := strconv.Atoi(v)
		intArr = append(intArr, num)
	}
	return intArr
}

func formatInput(input []string) [][]int {
	var result [][]int

	for i, v := range input {
		input[i] = strings.ReplaceAll(v, ":", "")
	}

	for _, v := range input {
		result = append(result, stringToIntArray(v))
	}
	return result
}

// change to possible values and return an array of possible values?
// func isPossibe(input []int) bool {
// 	length := len(input) - 2
// 	possibleValues := make([][]int, length)
// 	possibleValues[0] = append(possibleValues[0], (input[1] + input[2]))
// 	possibleValues[0] = append(possibleValues[0], (input[1] * input[2]))
// 	if len(input) > 3 {
// 		for i := 1; i < len(input)-2; i++ {
// 			for i1 := range possibleValues[i-1] {
// 				fmt.Println("input[i+3]: ", input[i+2])
// 				possibleValues[i] = append(possibleValues[i], (possibleValues[i-1][i1] + input[i+2]))
// 				possibleValues[i] = append(possibleValues[i], (possibleValues[i-1][i1] * input[i+2]))
// 			}
// 		}
// 	}
// 	fmt.Println(possibleValues)
// 	for _, v := range possibleValues[len(possibleValues)-1] {
// 		if v == input[0] {
// 			return true
// 		}
// 	}
// 	return false
// }

func possibleValues(input []int) []int {
	possibleValues := make([][]int, len(input)-1)

	// Add and multiply the first two values
	possibleValues[0] = append(possibleValues[0], (input[0] + input[1]))
	possibleValues[0] = append(possibleValues[0], (input[0] * input[1]))
	concat := strconv.Itoa(input[0]) + strconv.Itoa(input[1])
	concatAsInt, _ := strconv.Atoi(concat)
	possibleValues[0] = append(possibleValues[0], concatAsInt)

	for i := 1; i < len(input)-1; i++ {
		for j := range possibleValues[i-1] {
			possibleValues[i] = append(possibleValues[i], (possibleValues[i-1][j] + input[i+1]))
			possibleValues[i] = append(possibleValues[i], (possibleValues[i-1][j] * input[i+1]))

			concat := strconv.Itoa(possibleValues[i-1][j]) + strconv.Itoa(input[i+1])
			concatAsInt, _ := strconv.Atoi(concat)
			possibleValues[i] = append(possibleValues[i], concatAsInt)
		}
	}

	return possibleValues[len(possibleValues)-1]
}

func isPossible(input []int) bool {
	testValue, possibleValues := input[0], possibleValues(input[1:])
	return slices.Contains(possibleValues, testValue)
}

func part1(input [][]int) int {
	var total int
	for _, v := range input {
		if isPossible(v) {
			total += v[0]
		}
	}
	return total
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := formatInput(lines)
	fmt.Println(part1(input))
}
