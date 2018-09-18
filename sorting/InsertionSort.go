package sorting

import (
	"fmt"
)

func InsertionSorting(index int) {
	sortArray := getArray(index)
	for outer := 1; outer < index; outer++ {
		if sortArray[outer] < sortArray[outer-1] {
			seleted := sortArray[outer]
			inner := outer - 1
			for ; inner >= 0 && seleted < sortArray[inner]; inner-- {
				sortArray[inner+1] = sortArray[inner]
			}
			sortArray[inner+1] = seleted
		}
	}
	for i := 0; i < index; i++ {
		fmt.Println(sortArray[i])
	}
}
