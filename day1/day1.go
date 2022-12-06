package day1

import (
	"AdventOfCode2022/tools"
	"fmt"
	"sort"
	"strconv"
)

func Part1() {
	smallResult := calculateTopElfCalories(getInput("C:\\Repos\\AdventOfCode2022\\day1\\day1_ex.txt"))
	fmt.Println(smallResult)
	result := calculateTopElfCalories(getInput("C:\\Repos\\AdventOfCode2022\\day1\\day1.txt"))
	fmt.Println(result)
}

func Part2() {
	smallResult := calculateTopThreeElfCalories(getInput("C:\\Repos\\AdventOfCode2022\\day1\\day1_ex.txt"))
	fmt.Println(smallResult)
	result := calculateTopThreeElfCalories(getInput("C:\\Repos\\AdventOfCode2022\\day1\\day1.txt"))
	fmt.Println(result)
}

func calculateElfCalorieList(input []int) []int {
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
	return resultList
}

func getHighestValue(input []int) int {
	var largerNumber, temp int
	for _, element := range input {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	return largerNumber
}

func getHighestThreeValue(input []int) int {
	sort.Ints(input)
	sum := 0
	length := len(input)
	for i := 1; i < 4; i++ {
		sum = sum + input[length-i]
	}
	return sum
}

func calculateTopThreeElfCalories(input []int) int {
	calorieList := calculateElfCalorieList(input)
	return getHighestThreeValue(calorieList)
}

func calculateTopElfCalories(input []int) int {
	calorieList := calculateElfCalorieList(input)
	return getHighestValue(calorieList)
}

func getInput(filePath string) []int {
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
