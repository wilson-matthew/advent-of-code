package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(filePath string) (string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var rules string
	var updates string

	string := &rules

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			string = &updates
			continue
		}
		*string += scanner.Text() + " "
	}
	// trim whitespace appended to end of string
	rules = strings.TrimSpace(rules)
	updates = strings.TrimSpace(updates)
	return rules, updates, scanner.Err()

	// Change all of this function to just open the file and save to string var
	// then use strings.Split to get the two strings
}

func pagesInCorrectOrder(pages []string, rules string) bool {
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			tryRule := pages[j] + "|" + pages[i]
			if strings.Contains(rules, tryRule) {
				return false
			}
		}
	}
	return true
}

func correctOrderUpdates(updates string, rules string) [][]string {
	var correctOrderUpdates [][]string

	updatesArr := strings.Split(updates, " ")
	for _, v := range updatesArr {
		pages := strings.Split(v, ",")
		if pagesInCorrectOrder(pages, rules) {
			correctOrderUpdates = append(correctOrderUpdates, pages)
		}
	}
	return correctOrderUpdates
}

func incorrectOrderUpdates(updates string, rules string) [][]string {
	var incorrectOrderUpdates [][]string

	updatesArr := strings.Split(updates, " ")
	for _, v := range updatesArr {
		pages := strings.Split(v, ",")
		if !pagesInCorrectOrder(pages, rules) {
			incorrectOrderUpdates = append(incorrectOrderUpdates, pages)
		}
	}
	return incorrectOrderUpdates
}

func fixPageOrder(pages []string, rules string) []string {
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			tryRule := pages[j] + "|" + pages[i]
			if strings.Contains(rules, tryRule) {
				pages[i], pages[j] = pages[j], pages[i]
			}
		}
	}
	return pages
}

func fixIncorrectOrderUpdates(incorrectOrderUpdates [][]string, rules string) [][]string {
	var fixedOrderUpdates [][]string

	for _, v := range incorrectOrderUpdates {
		fixedOrderUpdates = append(fixedOrderUpdates, fixPageOrder(v, rules))
	}
	return fixedOrderUpdates
}

func sumOfMiddlePageNumbers(correctOrderUpdates [][]string) int {
	var total int

	for _, v := range correctOrderUpdates {
		pageNumber, _ := strconv.Atoi(v[len(v)/2])
		total += pageNumber
	}
	return total
}

func main() {
	rules, updates, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Sum of always correct middle page numbers: ", sumOfMiddlePageNumbers(correctOrderUpdates(updates, rules)))
	fmt.Println("Sum of fixed middle page numbers: ", sumOfMiddlePageNumbers(fixIncorrectOrderUpdates(incorrectOrderUpdates(updates, rules), rules)))
}
