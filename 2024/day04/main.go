package main

import (
	"fmt"
	"os"

	"github.com/dlclark/regexp2"
)

func openFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func countXmas(input string) int {
	right := regexp2.MustCompile(`(?=(XMAS))`, 0)
	left := regexp2.MustCompile(`(?=(SAMX))`, 0)

	// 140 is the line length of today's input
	down := regexp2.MustCompile(`(?s)(?=(X.{140}M.{140}A.{140}S))`, 0)
	up := regexp2.MustCompile(`(?s)(?=(S.{140}A.{140}M.{140}X))`, 0)

	downRight := regexp2.MustCompile(`(?s)(?=(X.{141}M.{141}A.{141}S))`, 0)
	downLeft := regexp2.MustCompile(`(?s)(?=(X.{139}M.{139}A.{139}S))`, 0)
	upRight := regexp2.MustCompile(`(?s)(?=(S.{139}A.{139}M.{139}X))`, 0)
	upLeft := regexp2.MustCompile(`(?s)(?=(S.{141}A.{141}M.{141}X))`, 0)

	matchedRight := len(regexp2FindAllString(right, input))
	matchedLeft := len(regexp2FindAllString(left, input))
	matchedDown := len(regexp2FindAllString(down, input))
	matchedUp := len(regexp2FindAllString(up, input))
	matchedDownRight := len(regexp2FindAllString(downRight, input))
	matchedDownLeft := len(regexp2FindAllString(downLeft, input))
	matchedUpRight := len(regexp2FindAllString(upRight, input))
	matchedUpLeft := len(regexp2FindAllString(upLeft, input))

	return matchedRight + matchedLeft + matchedDown + matchedUp + matchedDownRight + matchedDownLeft + matchedUpRight + matchedUpLeft
}

func countXMas(input string) int {
	// 140 is the line length of today's input
	mmass := regexp2.MustCompile(`(?s)(?=(M.{1}M.{139}A.{139}S.{1}S))`, 0)
	ssamm := regexp2.MustCompile(`(?s)(?=(S.{1}S.{139}A.{139}M.{1}M))`, 0)
	msams := regexp2.MustCompile(`(?s)(?=(M.{1}S.{139}A.{139}M.{1}S))`, 0)
	smasm := regexp2.MustCompile(`(?s)(?=(S.{1}M.{139}A.{139}S.{1}M))`, 0)

	matchedMmass := len(regexp2FindAllString(mmass, input))
	matchedSsamm := len(regexp2FindAllString(ssamm, input))
	matchedMsams := len(regexp2FindAllString(msams, input))
	matchedSmasm := len(regexp2FindAllString(smasm, input))

	return matchedMmass + matchedSsamm + matchedMsams + matchedSmasm
}

func main() {
	input, err := openFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Part 1: ", countXmas(input))
	fmt.Println("Part 2: ", countXMas(input))
}
