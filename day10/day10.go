package day10

import (
	"math"
	"sort"
)

type asteroidMap []string

type Position struct {
	x, y int
}

func MaxAsteroidsThatCanBeDetectedFromAStation(asteroids []string) (max int) {
	am := asteroidMap(asteroids)

	for y := 0; y < am.height(); y++ {
		for x := 0; x < am.width(); x++ {
			l := Position{x, y}
			if am.asteroidAt(l) {
				count := am.asteroidsVisibleFrom(l)
				if count > max {
					max = count
				}
			}
		}
	}

	return max
}

func (am *asteroidMap) asteroidAt(p Position) bool {
	return (*am)[p.y][p.x] == '#'
}

func (am *asteroidMap) width() int {
	return len((*am)[0])
}

func (am *asteroidMap) height() int {
	return len(*am)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func distanceTo(from, to Position) int {
	return abs(to.x-from.x) + abs(to.y-from.y)
}

func angleTo(from, to Position) float64 {
	return math.Mod(360-math.Atan2(float64(from.x-to.x), float64(from.y-to.y))*180.0/math.Pi, 360)
}

func (am *asteroidMap) positionsRelativeTo(p Position) map[float64][]Position {
	positions := make(map[float64][]Position)
	for toY := 0; toY < am.height(); toY++ {
		for toX := 0; toX < am.width(); toX++ {
			if p.x == toX && p.y == toY {
				continue
			}
			if am.asteroidAt(Position{toX, toY}) {
				angle := angleTo(p, Position{toX, toY})
				positions[angle] = append(positions[angle], Position{toX, toY})
			}
		}
	}
	return positions
}

func (am *asteroidMap) asteroidsVisibleFrom(p Position) (result int) {
	return len(am.positionsRelativeTo(p))
}

func VaporizedAsteroids(asteroids []string, giantLaser Position) (vaporized []Position) {
	am := asteroidMap(asteroids)

	positions := am.positionsRelativeTo(giantLaser)
	targetAngles := targetAngles(positions)

	for i := 0; i < am.totalAsteroids()-1; i++ {
		for _, angle := range targetAngles {
			targets := positions[angle]
			if len(targets) > 0 {
				target := closestTarget(targets, giantLaser)

				vaporized = append(vaporized, target)
				positions[angle] = locationsWithout(targets, target)
			}
		}
	}

	return vaporized
}

func targetAngles(positions map[float64][]Position) []float64 {
	angles := []float64{}
	for angle := range positions {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)

	return angles
}

func closestTarget(targets []Position, from Position) (target Position) {
	shortest := math.MaxInt64

	for _, position := range targets {
		distance := distanceTo(from, position)
		if distance < shortest {
			shortest = distance
			target = position
		}
	}

	return target
}

func locationsWithout(positions []Position, exclude Position) (result []Position) {
	for _, p := range positions {
		if p != exclude {
			result = append(result, p)
		}
	}
	return result
}

func (am *asteroidMap) totalAsteroids() (total int) {
	for y := 0; y < am.height(); y++ {
		for x := 0; x < am.width(); x++ {
			p := Position{x, y}
			if am.asteroidAt(p) {
				total++
			}
		}
	}
	return total
}
