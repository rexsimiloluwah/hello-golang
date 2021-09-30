// Implementation of common sorting algorithms in Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(BubbleSort([]float64{1, 5, 3, 2, 4, -9, 0, 12}, false))
}

func BubbleSort(arr []float64, reverse bool) []float64 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if reverse {
				if arr[j] < arr[j+1] {
					// Swapping
					temp := arr[j+1]
					arr[j+1] = arr[j]
					arr[j] = temp
				}
			} else {
				if arr[j] > arr[j+1] {
					// Swapping
					temp := arr[j+1]
					arr[j+1] = arr[j]
					arr[j] = temp
				}
			}
		}
	}
	return arr
}

// func InsertSort(arr []float64, reverse=bool) []float64{

// }
