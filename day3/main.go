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

	fmt.Printf("House Count: %d\n", countDeliveredHouses(strings.Split(string(input), "")))
}

func countDeliveredHouses(input []string) int {
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

	return places.uniquePositions()
}

func NewGrid() *grid {
	g := new(grid)
	g.positions = make([]point, 0)
	g.positions = append(g.positions, point{x: 0, y: 0})

	return g
}

type grid struct {
	positions []point
}

type point struct {
	x, y int
}

func (g grid) lastPosition() point {
	var lastPosition point

	if len(g.positions) > 0 {
		lastPosition = g.positions[len(g.positions)-1]
	} else {
		lastPosition = point{}
	}

	return lastPosition
}

func (g grid) uniquePositions() int {
	var seen []point

	for _, p := range g.positions {
		if !pointInSlice(seen, p) {
			seen = append(seen, p)
		}
	}

	return len(seen)
}

func (g *grid) right() {
	g.positions = append(g.positions, point{g.lastPosition().x + 1, g.lastPosition().y})
}

func (g *grid) left() {
	g.positions = append(g.positions, point{g.lastPosition().x - 1, g.lastPosition().y})
}

func (g *grid) up() {
	g.positions = append(g.positions, point{g.lastPosition().x, g.lastPosition().y + 1})
}

func (g *grid) down() {
	g.positions = append(g.positions, point{g.lastPosition().x, g.lastPosition().y - 1})
}

func pointInSlice(points []point, point point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}
