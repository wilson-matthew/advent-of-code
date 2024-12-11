package main

import (
	"fmt"
	"os"
	"slices"
)

func openFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	file, err := openFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var input []string

	for i := 0; i < len(file)-1; i++ {
		input = append(input, string(file[i]))
	}

	var lineLength int
	for i := 0; i < len(file)-1; i++ {
		lineLength += 1
		if string(file[i]) == "\n" {
			break
		}
	}

	var total int
	var alreadyCounted []int

	for i := 0; i < len(input)-1; i++ {
		if string(input[i]) != "." && string(input[i]) != "\n" && string(input[i]) != "#" {
			for j := i + 1; j < len(input)-1; j++ {
				if input[j] == input[i] {

					if !slices.Contains(alreadyCounted, i) {
						total += 1
						alreadyCounted = append(alreadyCounted, i)
					}
					if !slices.Contains(alreadyCounted, j) {
						total += 1
						alreadyCounted = append(alreadyCounted, j)
					}

					gap := j - i

					if (j % lineLength) > (i % lineLength) {
						// line is going left to right
						refPoint := i
						for refPoint-gap >= 0 && (refPoint%lineLength)-((j%lineLength)-(i%lineLength)) >= 0 {
							if !slices.Contains(alreadyCounted, refPoint-gap) {
								total += 1
								alreadyCounted = append(alreadyCounted, refPoint-gap)
								if input[refPoint-gap] == "." {
									input[refPoint-gap] = "#"
								}
							}
							refPoint -= gap
						}
						refPoint = i
						for refPoint+gap < len(input) && (refPoint%lineLength)+((j%lineLength)-(i%lineLength)) < lineLength-1 {
							if !slices.Contains(alreadyCounted, refPoint+gap) {
								total += 1
								alreadyCounted = append(alreadyCounted, refPoint+gap)
								if input[refPoint+gap] == "." {
									input[refPoint+gap] = "#"
								}
							}
							refPoint += gap
						}
					} else {
						// line is completely vertical or right to left
						refPoint := i
						for refPoint-gap >= 0 && (refPoint%lineLength)+((i%lineLength)-(j%lineLength)) < lineLength-1 {
							if !slices.Contains(alreadyCounted, refPoint-gap) {
								total += 1
								alreadyCounted = append(alreadyCounted, refPoint-gap)
								if input[refPoint-gap] == "." {
									input[refPoint-gap] = "#"
								}
							}
							refPoint -= gap
						}
						refPoint = i
						for refPoint+gap < len(input) && (refPoint%lineLength)-((i%lineLength)-(j%lineLength)) >= 0 {
							if !slices.Contains(alreadyCounted, refPoint+gap) {
								total += 1
								alreadyCounted = append(alreadyCounted, refPoint+gap)
								if input[refPoint+gap] == "." {
									input[refPoint+gap] = "#"
								}
							}
							refPoint += gap
						}
					}
				}
			}
		}
	}
	fmt.Println("Number of antenna: ", total)
}
