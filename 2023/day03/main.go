package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Confident this function works
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

// Pretty confident that this function is working fine
func numbersInLine(line string) []string {
	var numbersInLine []string
	var currentNumber string

	for _, value := range line {

		if unicode.IsDigit(value) {
			currentNumber += string(value)
		} else {
			if currentNumber != "" {
				numbersInLine = append(numbersInLine, currentNumber)
			}
			currentNumber = ""
		}
	}

	// Need this because if the last char in the line is a digit, then the last number won't get appended to numbersInLine
	if currentNumber != "" {
		numbersInLine = append(numbersInLine, currentNumber)
	}

	return numbersInLine
}

// change to take in filepath string, and call the readlines method within this method?
func sumOfPartNumbers(lines []string) int {
	var sumOfPartNumbers int

	// loop works fine
	for lineNumber, line := range lines {

		// looks fine
		numbersInLine := numbersInLine(line)

		// think this is fine
		minLineIndex := 0
		maxLineIndex := len([]rune(line)) - 1

		// fine
		lineBeforeLineNumber := lineNumber - 1
		lineAfterLineNumber := lineNumber + 1

		// fine
		// these will be set to 0
		var lineBefore string
		var lineAfter string

		// if these don't pass, then they'll still have the value of 0
		if lineBeforeLineNumber > 0 {
			lineBefore = lines[lineBeforeLineNumber]
		}

		if lineAfterLineNumber < len(lines) {
			lineAfter = lines[lineAfterLineNumber]
		}

		fmt.Println("line before line number:", lineBeforeLineNumber)
		fmt.Println("line number            :", lineNumber)
		fmt.Println("line after line number :", lineAfterLineNumber)
		fmt.Println()

		// Label break convention seems to be the same as variables, so MixedCaps or mixedCaps
		// All sample snippets from the Go language spec use MixedCaps for labels
	NumberLoop:
		for _, number := range numbersInLine {

			// use of index will cause issues if there is the same number twice in a line
			// becuase it will mean that the first instance of that number is always checked
			numberStartIndex := strings.Index(line, number)
			numberEndIndex := (numberStartIndex) + (len(number) - 1)

			for i := numberStartIndex - 1; i <= numberEndIndex+1; i++ {

				fmt.Println("i:", i)

				if i < minLineIndex || i > maxLineIndex {
					continue
				}

				if !strings.Contains("1234567890.", string(line[i])) {
					numberAsInt, _ := strconv.Atoi(number)
					sumOfPartNumbers += numberAsInt
					continue NumberLoop
				} else if (lineBeforeLineNumber >= 0) && (!strings.Contains("1234567890.", string(lineBefore[i]))) {
					numberAsInt, _ := strconv.Atoi(number)
					sumOfPartNumbers += numberAsInt
					continue NumberLoop
				} else if (lineAfterLineNumber < len(lines)) && (!strings.Contains("1234567890.", string(lineAfter[i]))) {
					numberAsInt, _ := strconv.Atoi(number)
					sumOfPartNumbers += numberAsInt
					continue NumberLoop
				}
			}
		}
	}

	return sumOfPartNumbers
}

func main() {
	lines, _ := readLines("./input.txt")
	fmt.Println("Sum of part numbers:", sumOfPartNumbers(lines))

	// for _, v := range lines {
	// 	fmt.Println("Numbers in line:", numbersInLine(v))
	// }
}
