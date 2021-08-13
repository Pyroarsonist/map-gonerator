package helpers

import (
	"math/rand"
	"time"
)

func GetRandomGenerator() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r
}

func MakeUnique(arr *[]int) {
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

func GetRandomItem(arr []int) int {
	r := GetRandomGenerator()
	return arr[r.Intn(len(arr))]
}
