package main

import (
	"strings"
)

type rediculousRules string

func (str rediculousRules) isNice() bool {
	if len(str) <= 2 {
		return false
	}

	return str.noBadStrings() && str.threeVowels() && str.twoConsecutiveLetters()
}

func (str rediculousRules) threeVowels() bool {
	var vowelCount int

	for _, letter := range strings.Split(string(str), "") {
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

func (str rediculousRules) twoConsecutiveLetters() bool {
	var previousLetter string

	for _, letter := range strings.Split(string(str), "") {
		if letter == previousLetter {
			return true
		} else {
			previousLetter = letter
		}
	}

	return false
}

func (str rediculousRules) noBadStrings() bool {
	for _, bad := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(string(str), bad) {
			return false
		}
	}

	return true
}
