package main

import "strings"

func main() {
	// call solution1 with testInput
	solution1(testInput)
	// call solution1 with realInput
	solution1(realInput)
	// call solution2 with testInput
	solution2(testInput)
	// call solution2 with realInput
	solution2(realInput)
}

// solution1 function takes a string and returns nothing
func solution1(s string) {

}

// solution2 function takes a string and returns nothing
func solution2(s string) {

}

// split string separated by two new lines
func splitStringByTwoNewLines(s string) []string {
	return strings.Split(s, "\n\n")
}

// split string by new line
func splitStringByNewLine(s string) []string {
	return strings.Split(s, "\n")
}

const (
	testInput = ``
	realInput = ``
)
