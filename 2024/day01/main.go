package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

func separateAndSortLists(lines []string) ([]string, []string) {
	var (
		list1 []string
		list2 []string
	)

	for _, v := range lines {
		splitLine := strings.Split(v, "   ")
		list1 = append(list1, splitLine[0])
		list2 = append(list2, splitLine[1])

	}
	slices.Sort(list1)
	slices.Sort(list2)
	return list1, list2
}

func differenceBetweenLists(list1 []string, list2 []string) int {
	var differenceBetweenLists int

	for i, v := range list1 {
		number1, _ := strconv.Atoi(v)
		number2, _ := strconv.Atoi(list2[i])
		difference := math.Abs(float64(number1 - number2))
		differenceBetweenLists += int(difference)
	}
	return differenceBetweenLists
}

func similarityScore(list1 []string, list2 []string) int {
	var similarityScore int

	for _, v1 := range list1 {
		var counter int
		for _, v2 := range list2 {
			if v1 == v2 {
				counter += 1
			}
		}
		v1int, _ := strconv.Atoi(v1)
		similarityScore += v1int * counter
	}
	return similarityScore
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	list1, list2 := separateAndSortLists(lines)
	fmt.Println(differenceBetweenLists(list1, list2))
	fmt.Println(similarityScore(list1, list2))
}
