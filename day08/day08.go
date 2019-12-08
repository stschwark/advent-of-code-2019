package day08

import (
	"fmt"
	"math"
)

func layerSize(width int, height int) int {
	return width * height
}

func numberOfLayers(image []int, width int, height int) int {
	return len(image) / layerSize(width, height)
}

func layersFromImage(image []int, width int, height int) (layers [][]int) {
	layerSize := layerSize(width, height)
	for i := 0; i < numberOfLayers(image, width, height); i++ {
		layers = append(layers, image[i*layerSize:(i+1)*layerSize])
	}
	return layers
}

func digitsInLayer(layer []int, digit int) (count int) {
	for i := range layer {
		if layer[i] == digit {
			count++
		}
	}
	return count
}

func checkSumLayer(layers [][]int) (checkSumLayer []int) {
	zerosInCheckSumLayer := math.MaxInt64
	for _, layer := range layers {
		if zeros := digitsInLayer(layer, 0); zeros < zerosInCheckSumLayer {
			zerosInCheckSumLayer = zeros
			checkSumLayer = layer
		}
	}
	return checkSumLayer
}

func CheckSum(image []int, width int, height int) (checksum int) {
	layers := layersFromImage(image, width, height)
	checkSumLayer := checkSumLayer(layers)
	return digitsInLayer(checkSumLayer, 1) * digitsInLayer(checkSumLayer, 2)
}

func visiblePixel(stackedPixels []int) int {
	for _, p := range stackedPixels {
		if p != 2 {
			return p
		}
	}
	return 2
}

func MergeLayers(image []int, width int, height int) []int {
	result := make([]int, width*height)
	layers := layersFromImage(image, width, height)
	for i := range result {
		pixels := make([]int, numberOfLayers(image, width, height))
		for z := range pixels {
			pixels[z] = layers[z][i]
		}
		result[i] = visiblePixel(pixels)
	}
	return result
}

func PrintImage(image []int, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := image[y*width+x]
			if pixel == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
