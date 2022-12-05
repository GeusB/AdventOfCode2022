package day2

import (
	//"AdventOfCode2022/tools"
	"fmt"
)

func Part1() {
	smallResult := CalulateScore(GetInput("C:\\Repos\\AdventOfCode2022\\day2\\day2_ex.txt"))
	fmt.Println(smallResult)
	result := CalulateScore(GetInput("C:\\Repos\\AdventOfCode2022\\day2\\day2.txt"))
	fmt.Println(result)
}

func Part2() {

}

func CalulateScore(input []int) int {
	// resultList := []int{}
	// length := len(input)
	sum := 0
	return sum
}

func GetInput(filePath string) []int {

	return nil
}
