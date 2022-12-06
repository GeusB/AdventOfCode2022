package day3

import (
	"AdventOfCode2022/tools"
	"fmt"
)

func Part1() {
	smallResult := CalulateScore("C:\\Repos\\AdventOfCode2022\\day3\\day3_ex.txt")
	fmt.Println(smallResult)
	result := CalulateScore("C:\\Repos\\AdventOfCode2022\\day3\\day3.txt")
	fmt.Println(result)
}

func Part2() {
}

func CalulateScore(filePath string) int {

	lines := tools.ReadFile(filePath)
	valueMap := makeCharValueMap()
	total := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineLength := len(line)
		firstPart := line[0 : lineLength/2]
		lastPart := line[lineLength/2 : lineLength]
		chars := []rune(firstPart)
		partLength := len(lastPart)

		target := makeMap(partLength, lastPart)
		match := getMatches(chars, target)

		lineSum := 0
		for j := 0; j < len(match); j++ {
			lineSum = lineSum + valueMap[match[j]]
		}

		total = total + lineSum
	}
	return total
}

func removeDuplicates(intSlice []string) []string {
	keys := make(map[string]bool)
	result := []string{}

	for _, x := range intSlice {
		if _, val := keys[x]; !val {
			keys[x] = true
			result = append(result, x)
		}
	}
	return result
}

func getMatches(chars []rune, target map[string]bool) []string {
	match := []string{}
	for _, val := range chars {
		str := string(val)
		if target[str] {
			match = append(match, str)
		}
	}
	return removeDuplicates(match)
}

func makeMap(partLength int, lastPart string) map[string]bool {
	target := make(map[string]bool)
	for j := 0; j < partLength; j++ {
		target[string(lastPart[j])] = true
	}
	return target
}

func makeCharValueMap() map[string]int {
	result := make(map[string]int)
	value := 1
	for j := 97; j < 123; j++ {
		result[string(j)] = value
		value++
	}
	for k := 65; k < 91; k++ {
		result[string(k)] = value
		value++
	}
	return result
}
