package utils

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

type Position struct {
	X, Y int
}

type Grid map[Position]int

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Grid) Print(paint func(Position, int) string) {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, -math.MaxInt64, -math.MaxInt64
	for p := range *g {
		minX = MinInt(minX, p.X)
		minY = MinInt(minY, p.Y)
		maxX = MaxInt(maxX, p.X)
		maxY = MaxInt(maxY, p.Y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			position := Position{x, y}
			val := (*g)[position]
			if paint != nil {
				fmt.Print(paint(position, val))
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}

func (g *Grid) Values() (values []int) {
	for key := range *g {
		values = append(values, (*g)[key])
	}
	return values
}
