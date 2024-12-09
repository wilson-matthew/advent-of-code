package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// It's returning the []string as nil because it couldn't open the file to load its contents, as well as the error produced
		return nil, err
	}
	// Defers the execution of a function until the surrounding function returns
	defer file.Close()

	// Creates a slice named lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func stringsToInts(lines []string) [][]int {
	var result [][]int

	for _, v := range lines {
		var ints []int
		for _, v := range strings.Split(v, " ") {
			int, _ := strconv.Atoi(v)
			ints = append(ints, int)
		}

		result = append(result, ints)

	}
	return result
}

func safe(line []int) bool {
	if (ascending(line) || descending(line)) && acceptableInterval(line) {
		return true
	}
	return false
}

func removeIndex(s []int, index int) []int {
	result := make([]int, 0, len(s)-1)
	result = append(result, s[:index]...)
	return append(result, s[index+1:]...)
}

func safeWithDampener(line []int) bool {
	for i := 0; i < len(line); i++ {
		modifiedReport := removeIndex(line, i)
		fmt.Println("modifiedReport: ", modifiedReport)
		if safe(modifiedReport) {
			return true
		}
	}
	return false
}

func ascending(line []int) bool {
	var biggest int
	for _, v := range line {
		if v < biggest {
			return false
		}
		biggest = v
	}
	return true
}

func descending(line []int) bool {
	smallest := line[0]
	for _, v := range line {
		if v > smallest {
			return false
		}
		smallest = v
	}
	return true
}

func acceptableInterval(line []int) bool {
	for i := 0; i < (len(line) - 1); i++ {
		difference := abs(line[i] - line[i+1])
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var safeReports int
	for _, v := range stringsToInts(lines) {
		if safe(v) {
			safeReports += 1
			fmt.Println(v, " safe")
		} else {
			fmt.Println(v, " unsafe")
		}
	}

	fmt.Println()

	var problemDampenedSafeReports int
	for _, v := range stringsToInts(lines) {
		if safeWithDampener(v) {
			problemDampenedSafeReports += 1
			fmt.Println(v, " safe")
		} else {
			fmt.Println(v, " unsafe")
		}
	}

	fmt.Println("No. of safe reports: ", safeReports)
	fmt.Println("No. of Problem Dampener safe reports: ", problemDampenedSafeReports)
}
