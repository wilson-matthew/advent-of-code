package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func retrieveGameId(line string) int {
	lineSplitOnComma := strings.Split(line, ":")
	// Game number is everything before the colon, i.e. 'Game 1'
	gameNumberString := lineSplitOnComma[0]
	// Split game number string on the space, so it becomes e.g. 'Game' and '24'
	// Then take just the second part, and convert it to an int
	gameNumberStringSplit := strings.Split(gameNumberString, " ")
	gameNumberString = gameNumberStringSplit[1]
	gameNumber, _ := strconv.Atoi(gameNumberString)
	return gameNumber
}

func retrieveRounds(line string) []string {
	lineSplitOnComma := strings.Split(line, ":")
	rounds := strings.Split(lineSplitOnComma[1], ";")
	return rounds
}

func numberOfDiceDrawnInRound(round string, colour string) int {
	var numberOfDiceDrawn int

	// Replace: returns a copy of round, where " " becomes "", with no limit on the number of replacements
	roundWithoutWhitespace := strings.Replace(round, " ", "", -1)
	drawnDice := strings.Split(roundWithoutWhitespace, ",")
	for _, value := range drawnDice {

		if strings.Contains(value, colour) {
			// Converts the characters between the start of the string and the colour to an int
			// I.e '17 blue' becomes an int '17'
			// Use blank identifier for error returned by Atoi method
			colourStartsAt := strings.Index(value, colour)
			numberOfDiceDrawn, _ = strconv.Atoi(value[0:colourStartsAt])
		}
	}

	return numberOfDiceDrawn
}

func maxNumberOfDiceDrawnInGame(rounds []string, colour string) int {
	var maxNumberOfDiceDrawn int

	for _, round := range rounds {
		numberOfDiceDrawnInRound := numberOfDiceDrawnInRound(round, colour)
		if numberOfDiceDrawnInRound > maxNumberOfDiceDrawn {
			maxNumberOfDiceDrawn = numberOfDiceDrawnInRound
		}
	}

	return maxNumberOfDiceDrawn
}

func gameIsPossible(rounds []string) bool {
	maxRedPossible := 12
	maxGreenPossible := 13
	maxBluePossible := 14

	maxRedInGame := maxNumberOfDiceDrawnInGame(rounds, "red")
	maxGreenInGame := maxNumberOfDiceDrawnInGame(rounds, "green")
	maxBlueInGame := maxNumberOfDiceDrawnInGame(rounds, "blue")

	if maxRedInGame <= maxRedPossible && maxGreenInGame <= maxGreenPossible && maxBlueInGame <= maxBluePossible {
		return true
	} else {
		return false
	}
}

func sumOfPossibleGameIds(allGames []string) int {
	var sumOfGameIds int

	for _, game := range allGames {

		rounds := retrieveRounds(game)

		if gameIsPossible(rounds) {
			sumOfGameIds += retrieveGameId(game)
		}
	}

	return sumOfGameIds
}

func main() {
	allGames, _ := readLines("./input.txt")
	solution := sumOfPossibleGameIds(allGames)
	fmt.Println(solution)
}
