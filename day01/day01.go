package day01

func CalculateFuelForMass(mass int) int {
	fuel := (mass / 3) - 2

	if fuel > 0 {
		return fuel
	}

	return 0
}

func CalculateBaseFuel(masses []int) int {
	total := 0

	for _, mass := range masses {
		total += CalculateFuelForMass(mass)
	}

	return total
}

func CalculateFuelPerModule(mass int) int {
	fuel := 0
	remainingMass := mass

	for remainingMass > 0 {
		additionalFuel := CalculateFuelForMass(remainingMass)

		fuel += additionalFuel
		remainingMass = additionalFuel
	}

	return fuel
}

func CalculateTotalFuelRequirements(masses []int) int {
	total := 0

	for _, mass := range masses {
		total += CalculateFuelPerModule(mass)
	}

	return total
}
