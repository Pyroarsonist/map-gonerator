package main

import (
	"math/rand"
)

func makeUnique(arr *[]int) {
	var uniqueArr []int
	keys := make(map[int]bool)
	for _, num := range *arr {
		if _, ok := keys[num]; !ok {
			keys[num] = true
			uniqueArr = append(uniqueArr, num)
		}
	}
	*arr = uniqueArr
}

func getRandom(arr []int) int {
	return arr[rand.Intn(len(arr))]
}
