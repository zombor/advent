package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Day 1 (part 1): %d\n", dayOne("input"))
	fmt.Printf("Day 1 (part 2): %d\n", dayOnePartTwo("input"))
}

func dayOne(file string) int {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var floor int

	for _, c := range strings.Split(string(contents), "") {
		if c == "(" {
			floor = floor + 1
		} else if c == ")" {
			floor = floor - 1
		}
	}

	return floor
}

func dayOnePartTwo(file string) int {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var floor int

	for i, c := range strings.Split(string(contents), "") {
		if c == "(" {
			floor = floor + 1
		} else if c == ")" {
			floor = floor - 1
		}

		if floor < 0 {
			return i
		}
	}

	return 0
}
