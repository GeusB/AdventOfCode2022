package day1

import (
	"AdventOfCode2022/tools"
	"fmt"
	"strconv"
)

func Part1() {
	smallResult := CalculateResult(SmallInput("C:\\Repos\\AdventOfCode2022\\day1\\day1_ex.txt"))
	fmt.Println(smallResult)
	result := CalculateResult(SmallInput("C:\\Repos\\AdventOfCode2022\\day1\\day1.txt"))
	fmt.Println(result)
}

func Part2() {
	smallResult := CalculateResult(SmallInput("C:\\Repos\\AdventOfCode2022\\day1\\day1_ex.txt"))
	fmt.Println(smallResult)
	result := CalculateResult(SmallInput("C:\\Repos\\AdventOfCode2022\\day1\\day1.txt"))
	fmt.Println(result)
}

func CalculateResult(input []int) int {
	resultList := []int{}
	length := len(input)
	sum := 0
	for i := 0; i < length; i++ {
		calories := input[i]
		sum = sum + calories
		if calories == 0 {
			resultList = append(resultList, sum)
			sum = 0
		}
	}
	if sum != 0 {
		resultList = append(resultList, sum)
		sum = 0
	}
	var largerNumber, temp int
	for _, element := range resultList {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	return largerNumber
}

func SmallInputLazy() []int {
	return []int{1000, 2000, 3000, 0, 4000, 0, 5000, 6000, 0, 7000, 8000, 9000, 0, 10000}
}

func SmallInput(filePath string) []int {
	lines := tools.ReadFile(filePath)
	length := len(lines)
	resultList := []int{}
	for i := 0; i < length; i++ {
		line := lines[i]
		number, _ := strconv.Atoi(line)
		resultList = append(resultList, number)
	}

	return resultList
}
