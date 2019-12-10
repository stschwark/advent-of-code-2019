package day10

import (
	"math"
	"sort"
)

type asteroidMap []string

type Location struct {
	x, y int
}

func MaxAsteroidsThatCanBeDetectedFromAStation(asteroids []string) (max int) {
	am := asteroidMap(asteroids)

	for y := 0; y < am.height(); y++ {
		for x := 0; x < am.width(); x++ {
			l := Location{x, y}
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

func (am *asteroidMap) asteroidAt(l Location) bool {
	return (*am)[l.y][l.x] == '#'
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

func distanceTo(from, to Location) int {
	return abs(to.x-from.x) + abs(to.y-from.y)
}

func angleTo(from, to Location) float64 {
	return math.Mod(360-math.Atan2(float64(from.x-to.x), float64(from.y-to.y))*180.0/math.Pi, 360)
}

func (am *asteroidMap) locationsRelativeTo(l Location) map[float64][]Location {
	locations := make(map[float64][]Location)
	for toY := 0; toY < am.height(); toY++ {
		for toX := 0; toX < am.width(); toX++ {
			if l.x == toX && l.y == toY {
				continue
			}
			if am.asteroidAt(Location{toX, toY}) {
				angle := angleTo(l, Location{toX, toY})
				l := locations[angle]
				locations[angle] = append(l, Location{toX, toY})
			}
		}
	}
	return locations
}

func (am *asteroidMap) asteroidsVisibleFrom(l Location) (result int) {
	return len(am.locationsRelativeTo(l))
}

func VaporizedAsteroids(asteroids []string, giantLaser Location) (vaporized []Location) {
	am := asteroidMap(asteroids)

	locations := am.locationsRelativeTo(giantLaser)
	targetAngles := targetAngles(locations)

	for i := 0; i < am.totalAsteroids()-1; i++ {
		for _, angle := range targetAngles {
			targets := locations[angle]
			if len(targets) > 0 {
				target := closestTarget(targets, giantLaser)

				vaporized = append(vaporized, target)
				locations[angle] = locationsWithout(targets, target)
			}
		}
	}

	return vaporized
}

func targetAngles(locations map[float64][]Location) []float64 {
	angles := []float64{}
	for angle := range locations {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)

	return angles
}

func closestTarget(targets []Location, from Location) (target Location) {
	shortest := math.MaxInt64

	for _, location := range targets {
		distance := distanceTo(from, location)
		if distance < shortest {
			shortest = distance
			target = location
		}
	}

	return target
}

func locationsWithout(locations []Location, exclude Location) (result []Location) {
	for _, l := range locations {
		if l != exclude {
			result = append(result, l)
		}
	}
	return result
}

func (am *asteroidMap) totalAsteroids() (total int) {
	for y := 0; y < am.height(); y++ {
		for x := 0; x < am.width(); x++ {
			l := Location{x, y}
			if am.asteroidAt(l) {
				total++
			}
		}
	}
	return total
}
