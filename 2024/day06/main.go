package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func stringToArray(s string) []string {
	var sArr []string
	for _, v := range s {
		sArr = append(sArr, string(v))
	}
	return sArr
}

func formatInput(input []string) [][]string {
	var result [][]string
	for _, v := range input {
		result = append(result, stringToArray(v))
	}
	return result
}

func startingPoints(input [][]string) (int, int) {
	for i, v := range input {
		currentLong := slices.Index(v, "^")
		if currentLong != -1 {
			currentLat := i
			return currentLat, currentLong
		}
	}
	return -1, -1
}

func changeDirection(currentDirection *string) {
	switch *currentDirection {
	case "up":
		*currentDirection = "right"
	case "right":
		*currentDirection = "down"
	case "down":
		*currentDirection = "left"
	case "left":
		*currentDirection = "up"
	}
}

func leavingMap(input [][]string, currentDirection string, currentLat int, currentLong int) bool {
	switch currentDirection {
	case "up":
		if currentLat == 0 {
			return true
		}
	case "right":
		if currentLong == len(input[currentLat])-1 {
			return true
		}
	case "down":
		if currentLat == len(input)-1 {
			return true
		}
	case "left":
		if currentLong == 0 {
			return true
		}
	}
	return false
}

func canMove(input [][]string, currentDirection string, currentLat int, currentLong int) bool {
	switch currentDirection {
	case "up":
		if input[currentLat-1][currentLong] != "#" {
			return true
		}
		return false
	case "right":
		if input[currentLat][currentLong+1] != "#" {
			return true
		}
		return false
	case "down":
		if input[currentLat+1][currentLong] != "#" {
			return true
		}
		return false
	case "left":
		if input[currentLat][currentLong-1] != "#" {
			return true
		}
		return false
	}
	return false
}

func move(input [][]string, currentDirection *string, currentLat *int, currentLong *int, distinctPositions *int) {
	switch *currentDirection {
	case "up":
		*currentLat -= 1
		if input[*currentLat][*currentLong] == "." {
			input[*currentLat][*currentLong] = "X"
			*distinctPositions += 1
		}
	case "right":
		*currentLong += 1
		if input[*currentLat][*currentLong] == "." {
			input[*currentLat][*currentLong] = "X"
			*distinctPositions += 1
		}
	case "down":
		*currentLat += 1
		if input[*currentLat][*currentLong] == "." {
			input[*currentLat][*currentLong] = "X"
			*distinctPositions += 1
		}
	case "left":
		*currentLong -= 1
		if input[*currentLat][*currentLong] == "." {
			input[*currentLat][*currentLong] = "X"
			*distinctPositions += 1
		}
	}
}

func distinctPositions(input [][]string) int {
	distinctPositions := 1
	currentLat, currentLong := startingPoints(input)
	currentDirection := "up"

	for {
		if leavingMap(input, currentDirection, currentLat, currentLong) {
			break
		} else if canMove(input, currentDirection, currentLat, currentLong) {
			move(input, &currentDirection, &currentLat, &currentLong, &distinctPositions)
		} else {
			changeDirection(&currentDirection)
		}
	}

	return distinctPositions
}

func guardInLoop(input [][]string) bool {
	distinctPositions := 1
	currentLat, currentLong := startingPoints(input)
	currentDirection := "up"
	var prevChangedLat []int
	var prevChangedLong []int

	for {
		if leavingMap(input, currentDirection, currentLat, currentLong) {
			break
		} else if canMove(input, currentDirection, currentLat, currentLong) {
			move(input, &currentDirection, &currentLat, &currentLong, &distinctPositions)
		} else {
			for !canMove(input, currentDirection, currentLat, currentLong) {
				changeDirection(&currentDirection)
			}

			prevChangedLat = append(prevChangedLat, currentLat)
			prevChangedLong = append(prevChangedLong, currentLong)

			if len(prevChangedLat) > 4 {
				for i := 0; i < len(prevChangedLat)-1; i++ {
					if currentLat == prevChangedLat[i] && currentLong == prevChangedLong[i] {
						return true
					}
				}
			}
		}
	}

	return false
}

func obstructionPositions(input [][]string) int {
	var total int
	for i, v := range input {
		for i1 := range v {
			if input[i][i1] == "X" {
				input[i][i1] = "#"

				if guardInLoop(input) {
					total += 1
				}

				input[i][i1] = "X"
			}
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

	distinctPositions := distinctPositions(input)
	fmt.Println("Distinct positions visited by guard: ", distinctPositions)

	obstructionPositions := obstructionPositions(input)
	fmt.Println("Number of obstruction positions to create loop: ", obstructionPositions)
}
