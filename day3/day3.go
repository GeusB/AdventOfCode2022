package day2

import (
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
	return 0
}
