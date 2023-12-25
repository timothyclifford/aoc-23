package main

import (
	"sort"
	"strconv"
	"unicode"
)

func main() {
	lines := readTextFile("./day01.txt")
	result := runDay01(lines)
	println(result)
}

func runDay01(lines []string) int {
	total := 0
	for _, line := range lines {
		words := filterLineForWords(line)
		digits := filterLineForDigits(line)
		joined := joinAndSortMaps(words, digits)
		if len(joined) == 0 {
			continue
		}

		first, last := getFirstLastValues(joined)
		numberString := strconv.Itoa(first) + strconv.Itoa(last)
		numberInt, _ := strconv.Atoi(numberString)

		total += numberInt
	}

	return total
}

func getFirstLastValues(values map[int]int) (first, last int) {
	keys := make([]int, 0)

	for k := range values {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for k, v := range keys {
		value := values[v]
		if k == 0 {
			first = value
		}
		last = value
	}

	return first, last
}

var numberWordsMap = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func filterLineForWords(line string) map[int]int {
	results := make(map[int]int)

	for k, v := range numberWordsMap {
		indices := findAllIndices(line, k)
		for _, vv := range indices {
			results[vv] = v
		}
	}

	return results
}

func filterLineForDigits(input string) map[int]int {
	result := make(map[int]int)
	chars := []rune(input)

	for k, v := range chars {
		if unicode.IsDigit(v) {
			theNumber, _ := strconv.Atoi(string(v))
			result[k] = theNumber
		}
	}

	return result
}
