package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliangruber/go-intersect"
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

func removeSpacesFromSlice(slice []string) []string {
	var sliceWithoutSpaces []string
	for _, v := range slice {
		if v != "" {
			sliceWithoutSpaces = append(sliceWithoutSpaces, v)
		}
	}
	return sliceWithoutSpaces
}

func winningNumbers(card string) []string {
	lineSplitOnComma := strings.Split(card, ":")
	winningNumbers := strings.Split(lineSplitOnComma[1], "|")
	winningNumbers = strings.Split(winningNumbers[0], " ")
	winningNumbers = removeSpacesFromSlice(winningNumbers)
	return winningNumbers
}

func cardNumbers(card string) []string {
	lineSplitOnComma := strings.Split(card, ":")
	cardNumbers := strings.Split(lineSplitOnComma[1], "|")
	cardNumbers = strings.Split(cardNumbers[1], " ")
	cardNumbers = removeSpacesFromSlice(cardNumbers)
	return cardNumbers
}

func countOfMatchingNumbers(card string) int {
	winningNumbers := winningNumbers(card)
	cardNumbers := cardNumbers(card)
	matchingNumbers := intersect.Simple(winningNumbers, cardNumbers)
	countOfMatchingNumbers := len(matchingNumbers)
	return countOfMatchingNumbers
}

func pointsInCard(card string) int {
	var pointsInCard int

	countOfMatchingNumbers := countOfMatchingNumbers(card)

	for i := 0; i <= countOfMatchingNumbers; i++ {
		if i == 0 {
			pointsInCard = 0
		} else if i == 1 {
			pointsInCard = 1
		} else {
			pointsInCard *= 2
		}
	}

	return pointsInCard
}

func pointsInAllCards(cards []string) int {
	var pointsInAllCards int

	for _, card := range cards {
		pointsInAllCards += pointsInCard(card)
	}

	return pointsInAllCards
}

func main() {
	lines, _ := readLines("./input.txt")
	pointsInAllCards := pointsInAllCards(lines)
	fmt.Println("Points in all cards:", pointsInAllCards)
}
