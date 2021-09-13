package main

import "fmt"

func main() {
	fizzBuzz(100)

	// Switch statements
	switch 6 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5, 6, 7, 8, 9, 10: // Using multiple tests
		fmt.Println("Greater than Four")
	default:
		fmt.Println("Default")
	}

	fizzBuzzSwitch(20)

	// Special switch case: - Type Switch

}

func fizzBuzz(num int) {
	// If-else if-else ladder statement
	for i := 0; i < num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzBuzzSwitch(num int) {
	for i := 0; i < num; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
