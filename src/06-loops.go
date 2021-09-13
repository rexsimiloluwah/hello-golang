package main

import "fmt"

func main() {
	// Simple loops (for loops)
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println(factorial(10))

	// Using break and continue (To print odd numbers less than 20)
	// d := 0
	// for d < 100 {
	// 	if d%2 == 0 {
	// 		d++
	// 		continue
	// 	}

	// 	if d > 20 {
	// 		break
	// 	}
	// 	fmt.Println(d)
	// 	d++
	// }

	// Infinite loop
	// for {
	// 	fmt.Printf("Hello")
	// }

	pyramidPattern(5)
	pyramidPattern2(5)

	// Working with collections in for loops
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, val := range s {
		fmt.Println(idx, val)
	}

	m := map[string]int{
		"Price":    1200,
		"Quantity": 5,
	}

	for _, val := range m {
		fmt.Println(val)
	}
}

func factorial(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result

	// Using recursion
	// if num <= 1 {
	// 	return num
	// }

	// return num * factorial(num-1)
}

// Nested loops :- Pyramid pattern
func pyramidPattern(height int) {
	for i := 1; i < height+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func pyramidPattern2(height int) {
	d := height
	for i := 1; i < (height * 2); i += 2 {
		d--
		for j := 0; j < d; j++ {
			fmt.Printf(" ")
		}

		for k := 0; k < i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
