package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func openFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func mulInstructions(input string) [][]string {
	ignore := regexp.MustCompile(`don't\(\).*?do\(\)`)
	capture := regexp.MustCompile(`mul\(([0-9]{1,3},[0-9]{1,3})\)`)

	cleaned := ignore.ReplaceAllString(input, "")
	match := capture.FindAllStringSubmatch(cleaned, -1)

	return match
}

func calculateInstructions(instructions [][]string) int {
	var total int
	for _, v := range instructions {
		// can get rid of this
		// instead add a capture group around each number in the capture regex
		// then num1 = v[1]
		// and num2 = v[2]
		numbers := strings.Split(v[1], ",")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		total += num1 * num2
	}
	return total
}

func main() {
	input, err := openFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mulInstructions := mulInstructions(input)
	fmt.Println("Answer: ", calculateInstructions(mulInstructions))
}
