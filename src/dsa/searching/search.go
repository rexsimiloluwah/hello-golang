package main

import "fmt"

func main() {
	fmt.Println(LinearSearch([]float64{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	fmt.Println(BinarySearchRecursively([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12, 0, 11))
}

func LinearSearch(arr []float64, el float64) int {
	i := 0
	for i < len(arr) {
		if arr[i] == el {
			return i
		}
		i++
	}
	return -1
}

func BinarySeachIteratively(arr []float64, el float64) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if el == arr[mid] {
			return mid
		} else if arr[mid] > el {
			high = mid - 1
		} else if arr[mid] < el {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursively(arr []float64, el float64, low int, high int) int {
	mid := low + (high-low)/2
	if low > high {
		return -1
	} else {
		if arr[mid] > el {
			return BinarySearchRecursively(arr, el, low, mid-1)
		} else if arr[mid] < el {
			return BinarySearchRecursively(arr, el, mid+1, high)
		} else {
			return mid
		}
	}
}
