package day1

import (
	"fmt"
)

func Part2() {

}

func Part1() {
	smallResult := CalculateResult(SmallInput())
	fmt.Println(smallResult)
}

func CalculateResult(input []int) int {
	resultList := []int{}
	length := len(input)
	sum := 0
	for i := 0; i < length-1; i++ {
		calories := input[i]
		sum = sum + calories
		if calories == 0 {
			resultList = append(resultList, sum)
			sum = 0
		}
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

func SmallInput() []int {
	return []int{1000, 2000, 3000, 0, 4000, 0, 5000, 6000, 0, 7000, 8000, 9000, 0, 10000}
}
