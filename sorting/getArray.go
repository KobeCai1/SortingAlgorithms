package sorting

import (
	"math/rand"
	"time"
)

func getArray(index int) []int32 {
	initArray := make([]int32, index+1)
	randomGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	for counter := 0; counter < index+1; counter++ {
		initArray[counter] = randomGenerator.Int31()
	}
	return initArray
}
