package sorting

import "fmt"

func MergeSort(index int) {
	sortArray := getArray(index)
	temp := make([]int32, index+1)
	fmt.Println(sortArray)
	splitArray(sortArray, 0, int32(index), temp)
	for i := 0; i < index+1; i++ {
		fmt.Println(sortArray[i])
	}
}

func splitArray(sArray []int32, first int32, last int32, temp []int32) {
	if first < last {
		mid := (first + last) / 2
		splitArray(sArray, first, mid, temp)
		splitArray(sArray, mid+1, last, temp)
		mergeArray(sArray, first, mid, last, temp)
	}
}

func mergeArray(mArray []int32, first int32, mid int32, last int32, temp []int32) {
	f, ml, mr, l := first, mid, mid+1, last
	var counter int32
	counter = 0
	for ml >= f && mr <= l {
		if mArray[f] <= mArray[mr] {
			temp[counter] = mArray[f]
			counter++
			f++
		} else {
			temp[counter] = mArray[mr]
			counter++
			mr++
		}
	}
	for ml >= f {
		temp[counter] = mArray[f]
		counter++
		f++
	}
	for mr <= l {
		temp[counter] = mArray[mr]
		counter++
		mr++
	}
	var i int32
	for i = 0; i < counter; i++ {
		mArray[first+i] = (temp[i])
	}
}
