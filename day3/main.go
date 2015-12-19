package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Santa House Count: %d\n",
		countDeliveredHouses(
			deliveredHouses(
				strings.Split(string(input), ""),
			).positions,
		),
	)

	fmt.Printf(
		"Santa+Robo Santa House Count: %d\n",
		countCombinedDeliveredHouses(strings.Split(string(input), "")),
	)
}

func deliveredHouses(input []string) grid {
	places := NewGrid()

	for _, direction := range input {
		if direction == ">" {
			places.right()
		} else if direction == "<" {
			places.left()
		} else if direction == "^" {
			places.up()
		} else if direction == "v" {
			places.down()
		}
	}

	return *places
}

func countDeliveredHouses(houses []point) int {
	return uniquePositions(houses)
}

func countCombinedDeliveredHouses(input []string) int {
	santaHouses := deliveredHouses(splitEven(input))
	roboSantaHouses := deliveredHouses(splitOdd(input))

	houses := santaHouses.positions

	for _, h := range roboSantaHouses.positions {
		houses = append(houses, h)
	}

	return countDeliveredHouses(houses)
}

func splitEven(strs []string) []string {
	output := make([]string, 0)

	for i, v := range strs {
		if i%2 == 0 {
			output = append(output, v)
		}
	}

	return output
}

func splitOdd(strs []string) []string {
	output := make([]string, 0)

	for i, v := range strs {
		if i%2 != 0 {
			output = append(output, v)
		}
	}

	return output
}

func uniquePositions(points []point) int {
	var seen []point

	for _, p := range points {
		if !pointInSlice(seen, p) {
			seen = append(seen, p)
		}
	}

	return len(seen)
}

func pointInSlice(points []point, point point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}
