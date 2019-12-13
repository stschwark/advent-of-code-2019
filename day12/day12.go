package day12

import (
	"reflect"
	"strconv"
)

const dimensions = 3

type Vector struct {
	p, v int
}

func Simulate(moons [][]Vector, steps int) (result [][]Vector) {
	for _, dvs := range transpose(moons) {
		after := simulateDimension(dvs, steps)
		result = append(result, after)
	}

	return transpose(result)
}

func transpose(matrix [][]Vector) (transposed [][]Vector) {
	for j := 0; j < len(matrix[0]); j++ {
		vs := []Vector{}
		for i := range matrix {
			vs = append(vs, matrix[i][j])
		}
		transposed = append(transposed, vs)
	}
	return transposed
}

func simulateDimension(vs []Vector, steps int) (result []Vector) {
	result = vs
	for i := 0; i < steps; i++ {
		result = applyGravity(result)
		result = applyVelocity(result)
	}
	return result
}

func gravity(a, b int) int {
	switch {
	case a > b:
		return -1
	case a < b:
		return 1
	default:
		return 0
	}
}

func applyGravity(vs []Vector) (result []Vector) {
	result = make([]Vector, len(vs))
	copy(result, vs)

	for i, v := range vs {
		for _, o := range vs {
			if !reflect.DeepEqual(v, o) {
				result[i].v += gravity(v.p, o.p)
			}
		}
	}

	return result
}

func applyVelocity(vs []Vector) (result []Vector) {
	for _, v := range vs {
		result = append(result, Vector{v.p + v.v, v.v})
	}
	return result
}

func CalculateTotalEnergy(moons [][]Vector) (totalEnergy int) {
	for _, moon := range moons {
		potentialEnergy := abs(moon[0].p) + abs(moon[1].p) + abs(moon[2].p)
		kenticEnergy := abs(moon[0].v) + abs(moon[1].v) + abs(moon[2].v)
		totalEnergy += potentialEnergy * kenticEnergy
	}
	return totalEnergy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateCycleCount(moons [][]Vector) (cycleCount int) {
	counts := []int{}
	for _, dvs := range transpose(moons) {
		counts = append(counts, cycleCountForDimension(dvs))
	}
	return lcm(counts...)
}

func cycleCountForDimension(dvs []Vector) int {
	cycleCount := 0
	current := dvs
	seen := make(map[string]int)
	for {
		cycleCount++

		current = applyGravity(current)
		current = applyVelocity(current)

		s := vectorsToString(current)
		if firstSeenAt, ok := seen[s]; ok {
			return cycleCount - firstSeenAt
		}

		seen[s] = cycleCount
	}
}

func vectorsToString(vs []Vector) (s string) {
	for _, v := range vs {
		s += strconv.Itoa(v.p) + "|" + strconv.Itoa(v.v) + ","
	}
	return s
}

func lcm(x ...int) int {
	if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	} else if x[0] == 0 && x[1] == 0 {
		return 0
	}
	return abs(x[0]*x[1]) / gcd(x[0], x[1])
}

func gcd(x ...int) int {
	if len(x) == 0 {
		return 0
	} else if len(x) == 2 {
		for x[1] != 0 {
			x[0], x[1] = x[1], x[0]%x[1]
		}
	} else if len(x) > 2 {
		return gcd(x[0], gcd(x[1:]...))
	}
	return abs(x[0])
}
