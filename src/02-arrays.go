package main

import (
	"fmt"
	"math/rand"
)

// Understanding Arrays and Slices
func main() {
	scores := [5]int{34, 45, 12, 67, 90}
	scoresFlexible := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("%v, %T\n", scores, scores)
	fmt.Printf("%v, %T\n", scoresFlexible, scoresFlexible)

	var students [4]string
	students[0] = "Okunowo Similoluwa"
	students[1] = "Salami Malcolm"
	students[2] = "David Adeleke"
	students[3] = "Okunowo Moyosoreoluwa (Naughty Girl)"
	fmt.Printf("%v, %T\n", students, students)
	fmt.Printf("Length of students array: - %v, %T\n", len(students), len(students))

	identityMatrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("%v, %T\n", identityMatrix, identityMatrix)

	// 3 x 3 Matrix
	var mat3 [3][3]int
	mat3[0] = [3]int{1, 2, 3}
	mat3[1] = [3]int{4, 5, 6}
	mat3[2] = [3]int{7, 8, 9}
	fmt.Println(mat3)

	// Copying of arrays in Go
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr3 := &arr1 // pointer to arr1
	arr2[2] = 9
	fmt.Println(arr2, arr1)
	arr3[2] = 9
	fmt.Println(*arr3, arr1)

	// Slicing arrays
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr4[2:6])
	fmt.Println(arr4[2:])

	// Slice - Lightweight structure that represents the portion of an array
	// A slice is not explicitly declared with the length in the square brackets
	// Capacity vs Length (cap vs len)
	// Make is used for initializing the slice
	arr5 := make([]int, 100) // Initialize an array with a length and capacity of 100
	//arr7 := make([]int, 10, 100) // Initialize an array with a length of 10 and capacity of 100
	fmt.Println(arr5)

	arr6 := []int{}
	arr6 = append(arr6, 5)
	fmt.Printf("%v\n", len(arr6))
	fmt.Printf("%v\n", cap(arr6))
	// Append an array using ... operator
	arr6 = append(arr6, []int{1, 2, 3, 4, 6}...)
	fmt.Printf("%v\n", arr6)

	multiArray := multiDimensionalArray(5, 5)
	fmt.Println(multiArray)
}

func multiDimensionalArray(nRow int, nCol int) [][]int {
	grid := make([][]int, nRow)

	for i := 0; i < nRow; i++ {
		grid[i] = make([]int, nCol)
		for j := 0; j < nCol; j++ {
			grid[i][j] = rand.Intn(10)
		}
	}

	return grid
}
