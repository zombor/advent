package main

import (
	"strings"
)

type betterRules string

func (str betterRules) isNice() bool {
	return str.hasSubpairTwice() && str.hasRepeatingNonConsecutiveLetters()
}

func (str betterRules) hasSubpairTwice() bool {
	stringSlice := str.split()

	for i, letter := range stringSlice {
		if i >= len(stringSlice)-1 {
			return false
		}

		pair := letter + stringSlice[i+1]

		if strings.Count(string(str), pair) >= 2 {
			return true
		}
	}

	return false
}

func (str betterRules) hasRepeatingNonConsecutiveLetters() bool {
	stringSlice := str.split()

	for i, letter := range stringSlice {
		if i >= len(stringSlice)-2 {
			return false
		} else if stringSlice[i+2] == letter {
			return true
		}
	}

	return false
}

func (str betterRules) split() []string {
	return strings.Split(string(str), "")
}
