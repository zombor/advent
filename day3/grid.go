package main

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
