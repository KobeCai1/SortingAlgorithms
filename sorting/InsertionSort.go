package sorting

import (
	"fmt"
)

func InsertionSorting(index int) {
	sortArray := getArray(index)
	for forward := 1; forward < index; forward++ {
		if sortArray[forward] < sortArray[forward-1] {
			seleted := sortArray[forward]
			backward := forward - 1
			for ; backward >= 0 && seleted < sortArray[backward]; backward-- {
				sortArray[backward+1] = sortArray[backward]
			}
			sortArray[backward+1] = seleted
		}
	}
	for i := 0; i < index; i++ {
		fmt.Println(sortArray[i])
	}
}
