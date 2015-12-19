package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var niceCount int

	for _, line := range strings.Split(string(contents), "\n") {
		if isNice(line) {
			niceCount++
		}
	}

	fmt.Printf("Nice Count: %d\n", niceCount)
}

func isNice(input string) bool {
	if len(input) <= 2 {
		return false
	}

	return noBadStrings(input) && threeVowels(input) && twoConsecutiveLetters(input)
}

func threeVowels(input string) bool {
	var vowelCount int

	for _, letter := range strings.Split(input, "") {
		for _, vowel := range []string{"a", "e", "i", "o", "u"} {
			if letter == vowel {
				vowelCount++
			}

			if vowelCount == 3 {
				return true
			}
		}
	}

	return false
}

func twoConsecutiveLetters(input string) bool {
	var previousLetter string

	for _, letter := range strings.Split(input, "") {
		if letter == previousLetter {
			return true
		} else {
			previousLetter = letter
		}
	}

	return false
}

func noBadStrings(input string) bool {
	for _, str := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(input, str) {
			return false
		}
	}

	return true
}
