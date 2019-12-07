package utils

import (
	"io/ioutil"
	"strings"
)

func ReadFromFile(name string) string {
	data, err := ioutil.ReadFile(name)

	if err != nil {
		panic("File reading error")
	}

	return strings.TrimSpace(string(data))
}

func ReadFromFileAndSplit(name string, delimiter string) []string {
	fileContent := ReadFromFile(name)

	return strings.Split(fileContent, delimiter)
}
