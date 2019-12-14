package day14

import (
	"math"
	"strconv"
	"strings"
)

type Reactant struct {
	quantity int
	chemical string
}

type Reaction struct {
	in  []Reactant
	out Reactant
}

func parseReactant(s string) Reactant {
	parts := strings.Split(s, " ")
	quantity, _ := strconv.Atoi(parts[0])
	return Reactant{quantity, parts[1]}
}

func parseReactants(ss []string) (r []Reactant) {
	for _, s := range ss {
		r = append(r, parseReactant(s))
	}
	return r
}

func parseReaction(s string) Reaction {
	parts := strings.Split(s, " => ")

	in := parseReactants(strings.Split(parts[0], ", "))
	out := parseReactant(parts[1])

	return Reaction{in, out}
}

func ParseRecipe(recipe []string) (reactions []Reaction) {
	for _, line := range recipe {
		reactions = append(reactions, parseReaction(line))
	}
	return reactions
}

type NanoFactory struct {
	reactions   []Reaction
	stock       map[string]int
	oreConsumed int
}

func nanoFactory(reactions []Reaction) (nf NanoFactory) {
	nf = NanoFactory{}
	nf.reactions = reactions
	nf.stock = make(map[string]int)
	return nf
}

func (nf *NanoFactory) reactionForChemical(chemical string) Reaction {
	for _, reaction := range nf.reactions {
		if reaction.out.chemical == chemical {
			return reaction
		}
	}
	panic("Cannot make " + chemical)
}

func (nf *NanoFactory) produce(quantity int, chemical string) {
	required := quantity

	// is the required chemical in stock?
	if stock, ok := nf.stock[chemical]; ok && stock > 0 {
		if required > stock {
			nf.stock[chemical] = 0
			required = required - stock
		} else {
			nf.stock[chemical] = stock - required
			required = 0
		}
	}

	// do we need to make more of the required chemical?
	if required > 0 {
		produced := 0

		if chemical == "ORE" {
			nf.oreConsumed += quantity
			produced = quantity
		} else {
			reaction := nf.reactionForChemical(chemical)

			// how many reactions should we run?
			numberOfReactions := int(math.Ceil(float64(required) / float64(reaction.out.quantity)))

			for _, input := range reaction.in {
				nf.produce(numberOfReactions*input.quantity, input.chemical)
			}

			produced = numberOfReactions * reaction.out.quantity
		}

		// have we produced more than we need?
		extra := produced - required

		nf.stock[chemical] = nf.stock[chemical] + extra
	}
}

func CalculateOreRequiredFor1Fuel(reactions []Reaction) int {
	nf := nanoFactory(reactions)

	nf.produce(1, "FUEL")

	return nf.oreConsumed
}

func CalculateFuelCreatedFromOre(reactions []Reaction, maxOrePerFuel int, ore int) (made int) {
	nf := nanoFactory(reactions)

	for {
		batchSize := (ore - nf.oreConsumed) / maxOrePerFuel

		if batchSize == 0 {
			batchSize = 1
		}

		nf.produce(batchSize, "FUEL")
		if nf.oreConsumed > ore {
			return made
		}

		made += batchSize
	}
}
