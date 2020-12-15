package main

import (
	"fmt"
)

func binaryFind(arr *[6]int, leftIndex, rightIndex, findVal int) {
	middle := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		fmt.Println("can't the value")
		return
	}

	if arr[middle] > findVal {
		binaryFind(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		binaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("the index is %d", middle)
	}

}

func main() {
	arr := [6]int{1, 5, 7, 8, 10, 22}
	binaryFind(&arr, 0, len(arr)-1, 7)
}
