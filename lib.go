package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readTextFile(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lines
}

func getKeys(theMap map[string]int) []string {
	keys := make([]string, 0, len(theMap))

	for key := range theMap {
		keys = append(keys, key)
	}

	return keys
}

func findAllIndices(mainString, subString string) []int {
	indices := []int{}
	index := -1

	for {
		index = strings.Index(mainString, subString)

		if index == -1 {
			break
		}

		indices = append(indices, index)
		mainString = mainString[index+len(subString):]
	}

	return indices
}

func joinMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key, value := range map2 {
		result[key] = value
	}

	return result
}

func joinAndSortMaps(map1, map2 map[int]int) map[int]int {
	keys := make([]int, 0, len(map1)+len(map2))
	for key := range map1 {
		keys = append(keys, key)
	}
	for key := range map2 {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	result := make(map[int]int)
	for _, key := range keys {
		if value, ok := map1[key]; ok {
			result[key] = value
		}
		if value, ok := map2[key]; ok {
			result[key] = value
		}
	}

	return result
}
