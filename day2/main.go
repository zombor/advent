package main

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"fmt"
)

func main() {
	var box = regexp.MustCompile(`^(?P<Length>\d+)x(?P<Width>\d+)x(?P<Height>\d+)$`)
	var totalSqFt, totalRibbonLength int

	contents, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(contents), "\n") {
		if line != "" {
			matches := box.FindStringSubmatch(line)

			l := parseString(matches[1])
			w := parseString(matches[2])
			h := parseString(matches[3])

			totalSqFt = totalSqFt + (2*l*w + 2*w*h + 2*h*l) + smallestSide(l, w, h)
			totalRibbonLength = totalRibbonLength + ribbonLength(l, w, h)
		}
	}

	fmt.Printf("Total SqFt: %d\n", totalSqFt)
	fmt.Printf("Total Ribbon Length: %d\n", totalRibbonLength)
}

func smallestSide(l, w, h int) int {
	side1 := l * w
	side2 := w * h
	side3 := l * h

	var min int = side1

	for _, i := range []int{side2, side3} {
		if i < min {
			min = i
		}
	}

	return min
}

func minLength(l, w, h int) int {
	sides := []int{l, w, h}
	sort.Sort(sort.IntSlice(sides))

	return sides[0]
}

func middleLength(l, w, h int) int {
	sides := []int{l, w, h}
	sort.Sort(sort.IntSlice(sides))

	return sides[1]
}

func ribbonLength(l, w, h int) int {
	smallest := minLength(l, w, h)
	middle := middleLength(l, w, h)

	return smallest + smallest + middle + middle + volume(l, w, h)
}

func volume(l, w, h int) int {
	return l * w * h
}

func parseString(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return i
}
