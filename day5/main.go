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

	var rediculousNiceCount int
	var betterNiceCount int

	for _, line := range strings.Split(string(contents), "\n") {
		if rediculousRules(line).isNice() {
			rediculousNiceCount++
		}
		if betterRules(line).isNice() {
			betterNiceCount++
		}
	}

	fmt.Printf("Rediculous Nice Count: %d\n", rediculousNiceCount)
	fmt.Printf("Better Nice Count    : %d\n", betterNiceCount)
}

type isNicer interface {
	isNice() bool
}
