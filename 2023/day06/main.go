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

func retrieveTimes() []string {
	lines, _ := readLines("./input.txt")
	timesWithSpaces := strings.Split((strings.Split(lines[0], ":")[1]), " ")
	var times []string
	for _, value := range timesWithSpaces {
		if value != "" {
			times = append(times, value)
		}
	}
	return times
}

func retrieveDistances() []string {
	lines, _ := readLines("./input.txt")
	timesWithSpaces := strings.Split((strings.Split(lines[1], ":")[1]), " ")
	var distances []string
	for _, value := range timesWithSpaces {
		if value != "" {
			distances = append(distances, value)
		}
	}
	return distances
}

func waysToBeatRecord(time string, distance string) int {
	timeAsInt, _ := strconv.Atoi(time)
	distanceAsInt, _ := strconv.Atoi(distance)

	var numberOfWaysToBeatRecord int

	for i := 0; i <= timeAsInt; i++ {
		if (i * (timeAsInt - i)) > distanceAsInt {
			numberOfWaysToBeatRecord++
		}
	}

	return numberOfWaysToBeatRecord
}

func partOneSolution() int {
	times := retrieveTimes()
	distances := retrieveDistances()

	var waysToBeatEachRecord []int
	var answer int

	for i, v := range times {
		waysToBeatEachRecord = append(waysToBeatEachRecord, waysToBeatRecord(v, distances[i]))
	}

	for i, v := range waysToBeatEachRecord {
		if i == 0 {
			answer = v
		} else {
			answer *= v
		}
	}

	return answer
}

func main() {
	fmt.Println("Part one solution:", partOneSolution())
	fmt.Println("Part two solution:", waysToBeatRecord("40817772", "219101213651089"))
}
