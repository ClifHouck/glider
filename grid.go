package main

import "fmt"

type Grid struct {
	dimension int
	grid      [][]bool
}

type Coordinate struct {
	X int
	Y int
}

func (g *Grid) Init(dimension int, liveCoordinates []Coordinate) {
	g.dimension = dimension

	g.grid = make([][]bool, g.dimension)
	for index := 0; index < g.dimension; index++ {
		g.grid[index] = make([]bool, g.dimension)
	}

	for _, c := range liveCoordinates {
		g.grid[c.X][c.Y] = true
	}
}

func (g *Grid) Dimension() int {
	return g.dimension
}

func (g *Grid) Print() {
	for index := 0; index < g.dimension; index++ {
		for inner_index := 0; inner_index < g.dimension; inner_index++ {
			if g.grid[index][inner_index] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Print("\n")
	}
}

func (g *Grid) IsValidCoordinate(x int, y int) bool {
	if x < 0 {
		return false
	} else if x >= g.dimension {
		return false
	} else if y < 0 {
		return false
	} else if y >= g.dimension {
		return false
	}
	return true
}

func (g *Grid) EnumerateValidNeighborCoordinates(x int, y int) []Coordinate {
	if !g.IsValidCoordinate(x, y) {
		return []Coordinate{}
	}

	neighbors := []Coordinate{
		Coordinate{X: x, Y: y + 1},
		Coordinate{X: x + 1, Y: y + 1},
		Coordinate{X: x + 1, Y: y},
		Coordinate{X: x + 1, Y: y - 1},
		Coordinate{X: x, Y: y - 1},
		Coordinate{X: x - 1, Y: y - 1},
		Coordinate{X: x - 1, Y: y},
		Coordinate{X: x - 1, Y: y + 1},
	}

	validNeighbors := []Coordinate{}
	for _, c := range neighbors {
		if g.IsValidCoordinate(c.X, c.Y) {
			validNeighbors = append(validNeighbors, c)
		}
	}
	return validNeighbors
}

func (g *Grid) CountLiveNeighbors(x int, y int) int {
	validNeighbors := g.EnumerateValidNeighborCoordinates(x, y)
	liveNeighbors := 0
	for _, coord := range validNeighbors {
		if g.grid[coord.X][coord.Y] {
			liveNeighbors++
		}
	}
	return liveNeighbors
}

func (g *Grid) NextCellState(x int, y int) bool {
	liveNeighbors := g.CountLiveNeighbors(x, y)
	// A live cell with 2 or 3 neighbors survives.
	if g.grid[x][y] && (liveNeighbors == 2 || liveNeighbors == 3) {
		return true
		// A dead cell with 3 live neighbors becomes alive.
	} else if !g.grid[x][y] && liveNeighbors == 3 {
		return true
	}
	// All other live cells die, or stay dead.
	return false
}

func (g *Grid) NextGrid() *Grid {
	nextGrid := Grid{}
	nextGrid.Init(g.dimension, []Coordinate{})

	for index := 0; index < g.dimension; index++ {
		for inner_index := 0; inner_index < g.dimension; inner_index++ {
			nextGrid.grid[index][inner_index] = g.NextCellState(index, inner_index)
		}
	}

	return &nextGrid
}

func (g *Grid) IsThereLife() bool {
	for x := 0; x < g.dimension; x++ {
		for y := 0; y < g.dimension; y++ {
			if g.grid[x][y] {
				return true
			}
		}
	}
	return false
}

func (g *Grid) IsEqual(other *Grid) bool {
	if g.dimension != other.dimension {
		return false
	}

	for x := 0; x < g.dimension; x++ {
		for y := 0; y < g.dimension; y++ {
			if g.grid[x][y] != other.grid[x][y] {
				return false
			}
		}
	}
	return true
}

func main() {
	g := &Grid{}
	g.Init(10, []Coordinate{
		Coordinate{X: 0, Y: 1},
		Coordinate{X: 1, Y: 2},
		Coordinate{X: 2, Y: 0},
		Coordinate{X: 2, Y: 1},
		Coordinate{X: 2, Y: 2},
	})

	for {
		g.Print()
		fmt.Println("=============================================")
		next := g.NextGrid()
		if g.IsEqual(next) {
			break
		}
		g = next
	}
}
