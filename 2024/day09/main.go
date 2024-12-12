package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func openFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func stringToIntArr(input string) ([]int, error) {
	var result []int
	for i := 0; i < len(input)-1; i++ {
		int, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return nil, err
		}
		result = append(result, int)
	}
	return result, nil
}

func insertAtIndex(slice []string, s string, index int) []string {
	slice = append(slice, " ")
	copy(slice[index+1:], slice[index:])
	slice[index] = s
	return slice
}

func expandInput(input []int) []string {
	var result []string
	var fileIndex int
	for i, v := range input {
		if i%2 == 0 {
			for j := 0; j < v; j++ {
				result = append(result, strconv.Itoa(fileIndex))
			}
			fileIndex += 1
		} else {
			for j := 0; j < v; j++ {
				result = append(result, ".")
			}
		}
	}
	return result
}

func moveFileBlocks(input []string) []string {
	var lastFile int

	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < len(input); j++ {
			if input[j] != "." {
				lastFile = j
			}
		}
		if slices.Index(input, ".") > lastFile {
			break
		}
		if input[i] != "." {
			firstFreeSpace := slices.Index(input, ".")
			input[firstFreeSpace], input[i] = input[i], input[firstFreeSpace]
		}
	}
	return input
}

func moveWholeFiles(input []string) []string {
	var currentFileId string
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != "." {
			currentFileId = input[i]
			break
		}
	}

	// Loops by decreasing file ID number
	for i, _ := strconv.Atoi(currentFileId); i >= 0; i-- {
		fileLength := (lastIndex(input, currentFileId) - slices.Index(input, currentFileId)) + 1

		var freeSpaceLength int
		var runningTotal int
		firstFreeSpace := slices.Index(input, ".")
		for j := firstFreeSpace; j < len(input); j++ {
			if input[j] == "." {
				runningTotal += 1
				if runningTotal >= fileLength {
					freeSpaceLength = runningTotal
					firstFreeSpace = j - freeSpaceLength + 1
					break
				}
			} else {
				runningTotal = 0
				if runningTotal > freeSpaceLength {
					freeSpaceLength = runningTotal
				}
			}
		}

		if freeSpaceLength >= fileLength {
			for i := 0; i < fileLength; i++ {
				input = slices.Replace(input, firstFreeSpace+i, firstFreeSpace+i+1, currentFileId)
				input = slices.Replace(input, lastIndex(input, currentFileId), lastIndex(input, currentFileId)+1, ".")
			}
		}

		currentFileId = strconv.Itoa(i - 1)
	}
	return input
}

func lastIndex[S ~[]E, E comparable](s S, v E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

func checksum(input []string) (int, error) {
	var total int

	for i, v := range input {
		if input[i] == "." {
			continue
		}
		fileID, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		total += i * fileID
	}
	return total, nil
}

func main() {
	file, err := openFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	intArr, err := stringToIntArr(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	expandedInput := expandInput(intArr)

	// compressed := moveFileBlocks(expandedInput)
	compressed := moveWholeFiles(expandedInput)

	checksum, err := checksum(compressed)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Checksum: ", checksum)
}
