package main

import (
	"fmt"
	"math"
	"strconv"
)

// Learning the very Basics of Golang (Variables, Data-types, Type Conversion, Operators)

// Global variables (You cannot use the : syntax here)
var price float32 = 35.45

var (
	name   string = "Similoluwa Okunowo"
	school string = "Obafemi Awolowo University, Ile-ife."
)

// iota starts at 0 for each const block and increments by one each time it is called in that same block
const (
	_  = iota             //iota = 0
	KB = 1 << (10 * iota) // iota = 1
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	var i int = 42
	// Using the : syntax, the compiler infers the variable type on runtime
	j := 50
	fmt.Println(i)
	var k float32 = 99
	fmt.Printf("%d, %T\n", j, j)
	fmt.Printf("%.2f, %T\n", k, k)
	fmt.Println(price)
	fmt.Printf("%s attends %s\n", name, school)

	// Type casting variables
	// Convert int to float64
	var num1 float64
	num1 = float64(j)
	fmt.Printf("%.2f, %T\n", num1, num1)
	// Convert string to int
	var str1 string
	str1 = strconv.Itoa(j) // Int to ASCII string
	fmt.Printf("%s, %T\n", str1, str1)

	// Primitive Data-types
	// Boolean
	var isRaining bool = false
	if isRaining {
		fmt.Println("Raining")
	} else {
		fmt.Println("Not Raining")
	}

	isEqual := 2 == 2
	fmt.Printf("%v, %T\n", isEqual, isEqual)

	// *Go assigns a default value of 0 to all declared variables

	// Numeric Data-types (Integer, Floating point data-types, double)
	var num2 uint16 = 12 // unsigned 16 bit integer
	fmt.Printf("%v, %T\n", num2, num2)
	fmt.Printf("%v, %T\n", float32(num2)/30, float32(num2)/30)

	// Mathematical operations cannot be performed on mismatched data-types (No implicit type conversion)

	// Bitwise Logical operators
	a := 10                 // 1010
	b := 12                 // 1100
	fmt.Printf("%v\n", a&b) // and (1000)
	fmt.Printf("%v\n", a|b) // or (1110)
	fmt.Printf("%v\n", a^b) // xor (0110)

	// Bit-shifting
	var d float32
	d = float32(a >> 2)
	fmt.Printf("%.2f\n", float32(d)) // a/4

	var num3 float64 = 23e20
	fmt.Printf("%.2f\n", num3)

	// Complex types (complex64 and complex128)
	var num4 complex64 = 1 + 2i // or complex(1,2)
	fmt.Printf("%v, %T\n", num4, num4)
	fmt.Printf("%v, %T\n", real(num4), real(num4))
	fmt.Printf("%v, %T\n", imag(num4), imag(num4))

	// Working with strings
	// Each character is a byte (uint8)
	s := "This is a string"
	sb := []byte(s)
	fmt.Printf("%v, %T\n", sb, sb)
	fmt.Printf("%v, %T\n", s[0], s[0])

	// Rune datatype, characters are uint32

	// Constants
	const myConst int = 34
	//myConst = 27 --> Error, cannot reassign a value to a constant
	fmt.Printf("%v, %T\n", myConst, myConst)
	const myConst2 = 23.5
	fmt.Printf("%v, %T\n", myConst2, myConst2)
	// Enumerated constants
	// Convert a number to human readable bytes using iota
	// Use PascalCase for exported constants and camelCase for the other
	var fileSize float64 = 400000000
	fmt.Printf("%.2fGB\n", float64(fileSize/GB))

	// Using the math module
	num8 := math.Pow(math.Sqrt(20), 2)
	fmt.Println(num8)
	fmt.Println(math.Ceil(12.5566))
}
