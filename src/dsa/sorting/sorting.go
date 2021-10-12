// Implementation of common sorting algorithms in Go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(BubbleSort([]float64{1, 5, 3, 2, 4, -9, 0, 12}, false))
	fmt.Println(SelectionSort([]float64{1, 5, 3, 2, 4, -9, 0, 12}, true))
	fmt.Println(InsertionSort([]float64{1, 5, 3, 2, 4, -9, 0, 12}))
	fmt.Println(ShellSort([]float64{1, 5, 3, 2, 4, -9, 0, 12}))
	fmt.Println(MergeSort([]float64{6, 0, 1, 5, 9, 8, 2}))
	//fmt.Println(partitionArray([]float64{6, 3, 1, 5, 9, 8, 2}))
	arr := []float64{6, 3, 1, 5, 9, 8, 2}
	QuickSort(arr)
	fmt.Println("Sorted array: ", arr)

}

// Selection Sort algorithm
// Average time complexity - 0(n^2)
func SelectionSort(arr []float64, reverse bool) []float64 {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if reverse {
				if arr[i] < arr[j] {
					temp := arr[j]
					arr[j] = arr[i]
					arr[i] = temp
				}
			} else {
				if arr[i] > arr[j] {
					temp := arr[j]
					arr[j] = arr[i]
					arr[i] = temp
				}
			}
		}
	}
	return arr
}

// Bubble Sort algorithm
// Adjacent elements are compared and swapped if they are not in order
// Average time complexity - 0(n^2)
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

// Insert Sort algorithm
// Insertion sort is very efficient when the size of the list is small or almost sorted
// Average time complexity - 0(n^2)
func InsertionSort(arr []float64) []float64 {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 {
			if temp < arr[j] {
				arr[j+1] = arr[j]
				j -= 1
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
	return arr
}

// Shell Sort algorithm : An optimized version of insertion sort which splits the original list into sublists
// and sorts the sublists via insertion sort.
// Average time complexity - 0(n^2)
func ShellSort(arr []float64) []float64 {
	var increment int
	// Increment for breaking the array into sub-arrays for optimized insertion sort
	increment = int(math.Floor(float64(len(arr) / 2)))
	for increment >= 1 {
		for i := increment; i < len(arr); i++ {
			// Insertion sort algorithm
			temp := arr[i]
			j := i - increment
			for j >= 0 {
				if temp < arr[j] {
					arr[j+increment] = arr[j]
					j = j - increment
				} else {
					break
				}
			}
			arr[j+increment] = temp
		}
		// Reduce the increment after each pass (The array becomes almost sorted after each pass)
		increment = int(math.Floor(float64(increment) / 2))
	}
	return arr
}

// Merge Sort algorithm
// Average time complexity - 0(n*log(n))
func MergeSort(arr []float64) []float64 {
	mid := len(arr) / 2
	if len(arr) < 2 {
		return arr
	}
	// Apply merge sort recursively on the left and right part of the array using divide and conquer
	left_arr := MergeSort(arr[:mid])
	right_arr := MergeSort(arr[mid:])
	return mergeHalves(left_arr, right_arr)
}

// Used for merging two sorted arrays
func mergeHalves(left_arr []float64, right_arr []float64) []float64 {
	// Merges two sorted arrays
	l_len := len(left_arr)                 // Length of the left arr
	r_len := len(right_arr)                // Length of the right arr
	result := make([]float64, l_len+r_len) // Slice for the final result

	i, j, k := 0, 0, 0
	for i < l_len && j < r_len {
		if left_arr[i] < right_arr[j] {
			result[k] = left_arr[i]
			i++
		} else {
			result[k] = right_arr[j]
			j++
		}
		k++
	}

	// For any remaining elements in the left arr
	for i < l_len {
		result[k] = left_arr[i]
		i++
		k++
	}

	// For any remaining elements in the right arr
	for j < r_len {
		result[k] = right_arr[j]
		j++
		k++
	}
	return result
}

func QuickSort(arr []float64) {
	if len(arr) <= 1 {
		return
	}
	pivotIdx := partitionArray(arr)
	fmt.Println(pivotIdx)
	QuickSort(arr[:pivotIdx])
	QuickSort(arr[pivotIdx:])
}

// A utility function for partitioning an array
// The first element is used as the pivot, and the elements in the array less than the pivot are placed on the left side
// elements in the array greater than the pivot are placed on the right side
// The final index of the pivot in the array is returned
func partitionArray(arr []float64) int {
	pivot := arr[len(arr)-1] // select first element as the pivot
	i, j := -1, 0
	for j < len(arr) {
		if pivot > arr[j] {
			i += 1
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
		j++
	}
	arr[len(arr)-1], arr[i+1] = arr[i+1], pivot
	fmt.Println(arr)
	return (i + 1)
}
