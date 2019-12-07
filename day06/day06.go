package day06

import (
	"math"
	"regexp"
)

type spaceMap struct {
	orbits map[string]string
}

func (sm *spaceMap) orbitFor(object string) (string, bool) {
	val, ok := sm.orbits[object]
	return val, ok
}

func (sm *spaceMap) countOrbitsTo(outer string, inner string) int {
	direct, isInOrbit := sm.orbitFor(outer)

	if !isInOrbit {
		return 0
	}

	if direct == inner {
		return 1
	}

	directAnIndirectOrbits := sm.countOrbitsTo(direct, inner)
	if directAnIndirectOrbits > 0 {
		return directAnIndirectOrbits + 1
	}

	return 0
}

func createSpaceMap(orbits []string) spaceMap {
	orbitsRegex := regexp.MustCompile(`(.+)\)(.+)`)
	inOrbitOf := make(map[string]string)
	for _, orbit := range orbits {
		match := orbitsRegex.FindAllStringSubmatch(orbit, 2)
		center := match[0][1]
		object := match[0][2]

		inOrbitOf[object] = center
	}

	return spaceMap{inOrbitOf}
}

func CalculateNumberOfDirectAndIndirectOrbits(orbits []string) int {
	sm := createSpaceMap(orbits)

	count := 0
	for name := range sm.orbits {
		count += sm.countOrbitsTo(name, "COM")
	}

	return count
}

func CalculateMinimumNumberOfOrbitalTransfers(orbits []string, a string, b string) int {
	sm := createSpaceMap(orbits)

	result := math.MaxInt32
	for name := range sm.orbits {
		orbitsToA := sm.countOrbitsTo(a, name)
		orbitsToB := sm.countOrbitsTo(b, name)
		distance := orbitsToA + orbitsToB

		if orbitsToA != 0 && orbitsToB != 0 && distance < result {
			result = distance
		}
	}

	return result - 2
}
