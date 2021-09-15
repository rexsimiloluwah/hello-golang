package main

import (
	"errors"
	"fmt"
	"math/cmplx"
)

func main() {
	root1, root2, err := quadSolver(1, 2, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(root1, root2)

	name := "Similoluwa"
	fmt.Println(name)
	mutateName(&name)
	fmt.Println(name)

	avg := average(1, 2, 3, 4, 5)
	fmt.Println("The average of 1,2,3,4,5 is: - ", avg)

	// Anonymous function
	func() {
		fmt.Println("Anonymous function which is also immediately invoked.")
	}()

	divide := func(a float64, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("Zero Division error")
		}
		return a / b, nil
	}

	d, err := divide(5, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	// Methods
	g1 := greeter{
		name:     "Similoluwa Okunowo",
		greeting: "Hello",
	}

	g1.greet()

}

func quadSolver(x1 float64, x2 float64, x3 float64) (complex128, complex128, error) {
	if x1 <= 0 {
		return 0, 0, errors.New("a > 0 for a valid Quadratic equation.")
	}

	a, b, c := complex(x1, 0), complex(x2, 0), complex(x3, 0)
	discriminant := cmplx.Pow(b, 2) - 4*a*c
	r1 := (-b + cmplx.Sqrt(discriminant)) / (2 * a)
	r2 := (-b - cmplx.Sqrt(discriminant)) / (2 * a)
	return r1, r2, nil
}

/**
* Function with Pointers
 */
func mutateName(name *string) {
	*name = "[MUTATED] Similoluwa"
}

/**
* Working with Variadic functions - This function accepts a variable number of arguments
 */
func average(values ...int) (result int) {
	fmt.Println(values)
	count, sum := 0, 0
	// The values collection will be a slice
	for _, v := range values {
		count++
		sum += v
	}
	result = sum / count
	// result is implicitly returned
	return
}

/**
* Working with methods on different types (i.e. struct)
 */

type greeter struct {
	name     string
	greeting string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, ",", g.name)
}
