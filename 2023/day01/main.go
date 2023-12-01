package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func findDigitsInString(word string) []rune {
	var digitsInWord []rune

	// This returns byte position of the character and the character
	for _, char := range word {
		if unicode.IsDigit(char) {
			digitsInWord = append(digitsInWord, char)
		}
	}

	return digitsInWord
}

func constructCalibrationValue(digitsInWord []rune) int {
	firstDigitRune := digitsInWord[0]
	lastDigitRune := digitsInWord[len(digitsInWord)-1]
	firstDigitString := string(firstDigitRune)
	lastDigitString := string(lastDigitRune)
	calibrationValueAsString := firstDigitString + lastDigitString
	calibrationValueAsInt, _ := strconv.Atoi(calibrationValueAsString)
	return calibrationValueAsInt
}

func recoverCalibrationValues(lines []string) int {
	var calibrationValueTotal int

	// Two values are returned when ranging over a slice: index, and a copy of the element at that index
	for _, word := range lines {
		digitsInWord := findDigitsInString(word)
		calibrationValue := constructCalibrationValue(digitsInWord)
		calibrationValueTotal += calibrationValue
	}

	return calibrationValueTotal
}

func main() {
	lines, _ := readLines("./input.txt")
	fmt.Println(recoverCalibrationValues(lines))
}
