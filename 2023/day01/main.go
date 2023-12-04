package main

import (
	"bufio"
	"fmt"
	"math"
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

// Function to find first occuring number
// Iterate over the digits slice
// Use strings.index(str, substr) to check whether that digit exists, and if it does, at what index
// Check if digit exists in string
// Check if it occurs at an earlier index than any of the others
// If it doesn't, just ignore it
// If it does, set var firstOccurenceOfDigit to that digit

func findFirstDigitInString(str string) string {
	var firstDigitInString string
	currentDigitIndex := math.MaxInt16
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, digit := range digits {
		indexOfDigit := strings.Index(str, digit)
		if indexOfDigit > -1 && indexOfDigit < currentDigitIndex {
			firstDigitInString = digit
			currentDigitIndex = indexOfDigit
		}
	}
	return firstDigitInString
}

func findLastDigitInString(str string) string {
	var lastDigitInString string
	currentDigitIndex := -1
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, digit := range digits {
		indexOfDigit := strings.LastIndex(str, digit)
		if indexOfDigit > -1 && indexOfDigit > currentDigitIndex {
			lastDigitInString = digit
			currentDigitIndex = indexOfDigit
		}
	}
	return lastDigitInString
}

func convertNumbersToDigits(digit string) string {
	var digitAsInt string

	switch digit {
	case "1", "one":
		digitAsInt = "1"
	case "2", "two":
		digitAsInt = "2"
	case "3", "three":
		digitAsInt = "3"
	case "4", "four":
		digitAsInt = "4"
	case "5", "five":
		digitAsInt = "5"
	case "6", "six":
		digitAsInt = "6"
	case "7", "seven":
		digitAsInt = "7"
	case "8", "eight":
		digitAsInt = "8"
	case "9", "nine":
		digitAsInt = "9"
	}

	return digitAsInt
}

func constructCalibrationValue(firstDigit string, lastDigit string) int {
	calibrationValueAsString := firstDigit + lastDigit
	calibrationValueAsInt, _ := strconv.Atoi(calibrationValueAsString)
	return calibrationValueAsInt
}

func recoverCalibrationValues(lines []string) int {
	var calibrationValueTotal int

	// Two values are returned when ranging over a slice: index, and a copy of the element at that index
	for _, word := range lines {
		firstDigitInString := findFirstDigitInString(word)
		lastDigitInString := findLastDigitInString(word)
		firstDigitAsDigit := convertNumbersToDigits(firstDigitInString)
		lastDigitAsDigit := convertNumbersToDigits(lastDigitInString)
		calibrationValue := constructCalibrationValue(firstDigitAsDigit, lastDigitAsDigit)
		calibrationValueTotal += calibrationValue
	}

	return calibrationValueTotal
}

func main() {
	lines, _ := readLines("./input.txt")
	fmt.Println(recoverCalibrationValues(lines))
}
