package day3

import (
	"AdventOfCode2022/tools"
	"fmt"
)

func Part1() {
	smallResult := calculateCompartmentsScore("C:\\Repos\\AdventOfCode2022\\day3\\day3_ex.txt")
	fmt.Println(smallResult)
	result := calculateCompartmentsScore("C:\\Repos\\AdventOfCode2022\\day3\\day3.txt")
	fmt.Println(result)
}

func Part2() {
	smallResult := calculateGroupScore("C:\\Repos\\AdventOfCode2022\\day3\\day3_ex.txt")
	fmt.Println(smallResult)
	result := calculateGroupScore("C:\\Repos\\AdventOfCode2022\\day3\\day3.txt")
	fmt.Println(result)
}

func calculateGroupScore(filePath string) int {

	lines := tools.ReadFile(filePath)
	valueMap := makeCharValueMap()
	total := 0
	groups := tools.ChunkSlice(lines, 3)
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		firstLine := []rune(group[0])
		secondLine := makeMap(group[1])
		thirdLine := makeMap(group[2])
		preMatch := getMatches(firstLine, secondLine)
		match := getMatches2(preMatch, thirdLine)
		total = total + valueMap[match[0]]
	}
	return total
}

func calculateCompartmentsScore(filePath string) int {

	lines := tools.ReadFile(filePath)
	valueMap := makeCharValueMap()
	total := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineLength := len(line)
		firstPart := line[0 : lineLength/2]
		lastPart := line[lineLength/2 : lineLength]
		chars := []rune(firstPart)
		target := makeMap(lastPart)
		match := getMatches(chars, target)

		lineSum := 0
		for j := 0; j < len(match); j++ {
			lineSum = lineSum + valueMap[match[j]]
		}

		total = total + lineSum
	}
	return total
}

func getMatches(chars []rune, target map[string]bool) []string {
	match := []string{}
	for _, val := range chars {
		str := string(val)
		if target[str] {
			match = append(match, str)
		}
	}
	return tools.RemoveDuplicates(match)
}

func getMatches2(chars []string, target map[string]bool) []string {
	match := []string{}
	for _, val := range chars {
		str := string(val)
		if target[str] {
			match = append(match, str)
		}
	}
	return tools.RemoveDuplicates(match)
}

func makeMap(lastPart string) map[string]bool {
	partLength := len(lastPart)
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
		result[string(rune(j))] = value
		value++
	}
	for k := 65; k < 91; k++ {
		result[string(rune(k))] = value
		value++
	}
	return result
}
