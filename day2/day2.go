package day2

import (
	"AdventOfCode2022/tools"
	"fmt"
)

func Part1() {
	smallResult := CalulateScore("C:\\Repos\\AdventOfCode2022\\day2\\day2_ex.txt")
	fmt.Println(smallResult)
	result := CalulateScore("C:\\Repos\\AdventOfCode2022\\day2\\day2.txt")
	fmt.Println(result)
}

func Part2() {

}

func CalulateScore(filePath string) int {
	lines := tools.ReadFile(filePath)
	length := len(lines)
	total := 0
	for i := 0; i < length; i++ {
		line := lines[i]
		result := 0
		switch line {
		case "A X":
			result = 4
		case "A Y":
			result = 8
		case "A Z":
			result = 3
		case "B X":
			result = 1
		case "B Y":
			result = 5
		case "B Z":
			result = 9
		case "C X":
			result = 7
		case "C Y":
			result = 2
		case "C Z":
			result = 6
		}
		total = total + result
	}
	return total
}
