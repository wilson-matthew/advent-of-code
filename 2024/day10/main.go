package main

import (
	"fmt"
	"log"
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

func lineLength(s string) int {
	var lineLength int
	for i := 0; i < len(s); i++ {
		lineLength += 1
		if string(s[i]) == "\n" {
			break
		}
	}
	return lineLength
}

func trailheadScore(input string, currentPos int, lineLength int, currentHeight int, total *int, alreadyCounted *[]int) {
	// directly up from current position
	if currentPos-lineLength >= 0 && string(input[currentPos-lineLength]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			if !slices.Contains(*alreadyCounted, currentPos-lineLength) {
				*total += 1
				*alreadyCounted = append(*alreadyCounted, currentPos-lineLength)
			}
		}
		trailheadScore(input, currentPos-lineLength, lineLength, currentHeight+1, total, alreadyCounted)
	}
	// directly below current position
	if currentPos+lineLength < len(input) && string(input[currentPos+lineLength]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			if !slices.Contains(*alreadyCounted, currentPos+lineLength) {
				*total += 1
				*alreadyCounted = append(*alreadyCounted, currentPos+lineLength)
			}
		}
		trailheadScore(input, currentPos+lineLength, lineLength, currentHeight+1, total, alreadyCounted)
	}
	// left
	if currentPos-1 >= 0 && string(input[currentPos-1]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			if !slices.Contains(*alreadyCounted, currentPos-1) {
				*total += 1
				*alreadyCounted = append(*alreadyCounted, currentPos-1)
			}
		}
		trailheadScore(input, currentPos-1, lineLength, currentHeight+1, total, alreadyCounted)
	}
	// right
	if currentPos+1 < len(input) && string(input[currentPos+1]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			if !slices.Contains(*alreadyCounted, currentPos+1) {
				*total += 1
				*alreadyCounted = append(*alreadyCounted, currentPos+1)
			}
		}
		trailheadScore(input, currentPos+1, lineLength, currentHeight+1, total, alreadyCounted)
	}
}

// change so it works with any slice type
func containsSlice(s1 [][]int, s2 []int) bool {
out:
	for _, v := range s1 {
		for i1, v2 := range v {
			if v2 != s2[i1] {
				continue out
			}
		}
		return true
	}
	return false
}

func trailheadRating(input string, currentPos int, lineLength int, currentHeight int, route *[]int, successfulRoutes *[][]int) {
	*route = append(*route, currentPos)
	// directly up from current position
	if currentPos-lineLength >= 0 && string(input[currentPos-lineLength]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			*route = append(*route, currentPos-lineLength)
			if !containsSlice(*successfulRoutes, *route) {
				*successfulRoutes = append(*successfulRoutes, *route)
			}
			*route = nil
		}
		trailheadRating(input, currentPos-lineLength, lineLength, currentHeight+1, route, successfulRoutes)
	}
	// directly below current position
	if currentPos+lineLength < len(input) && string(input[currentPos+lineLength]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			*route = append(*route, currentPos+lineLength)
			if !containsSlice(*successfulRoutes, *route) {
				*successfulRoutes = append(*successfulRoutes, *route)
			}
			*route = nil
		}
		trailheadRating(input, currentPos+lineLength, lineLength, currentHeight+1, route, successfulRoutes)
	}
	// left
	if currentPos-1 >= 0 && string(input[currentPos-1]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			*route = append(*route, currentPos-1)
			if !containsSlice(*successfulRoutes, *route) {
				*successfulRoutes = append(*successfulRoutes, *route)
			}
			*route = nil
		}
		trailheadRating(input, currentPos-1, lineLength, currentHeight+1, route, successfulRoutes)
	}
	// right
	if currentPos+1 < len(input) && string(input[currentPos+1]) == strconv.Itoa(currentHeight+1) {
		if currentHeight+1 == 9 {
			*route = append(*route, currentPos+1)
			if !containsSlice(*successfulRoutes, *route) {
				*successfulRoutes = append(*successfulRoutes, *route)
			}
			*route = nil
		}
		trailheadRating(input, currentPos+1, lineLength, currentHeight+1, route, successfulRoutes)
	}
}

func sumOfTrailheadScores(input string) int {
	lineLength := lineLength(input)
	var total int
	var score int

	for i := 0; i < len(input); i++ {
		if string(input[i]) == "0" {
			currentPos := i
			currentHeight := 0
			score = 0
			var alreadyCounted []int
			trailheadScore(input, currentPos, lineLength, currentHeight, &score, &alreadyCounted)
			// fmt.Println("trailheadScore: ", score)
			// fmt.Println(alreadyCounted)
			total += score
		}
	}
	return total
}

func sumOfTrailheadRatings(input string) int {
	lineLength := lineLength(input)
	var total int

	for i := 0; i < len(input); i++ {
		if string(input[i]) == "0" {
			currentPos := i
			currentHeight := 0
			var route []int
			var successfulRoutes [][]int
			trailheadRating(input, currentPos, lineLength, currentHeight, &route, &successfulRoutes)
			// fmt.Println("trailheadScore: ", score)
			// fmt.Println(alreadyCounted)
			fmt.Println(successfulRoutes)
			total += len(successfulRoutes)
		}
	}
	return total
}

func main() {
	input, err := openFile("test-input-4.txt")
	if err != nil {
		log.Fatal(err)
	}
	sumOfTrailheadScores := sumOfTrailheadScores(input)
	fmt.Println("sumOfTrailheadScores: ", sumOfTrailheadScores)
	fmt.Println("sumOfTrailheadRatings: ", sumOfTrailheadRatings(input))
}
