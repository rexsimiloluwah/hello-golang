package main

import "fmt"

type myStruct struct {
	id   int
	name string
}

func main() {
	a := 20
	var b *int = &a
	fmt.Printf("Address of a:- %p\n", b)
	*b = 40 // de-referencing the pointer using *b
	fmt.Println(a)

	// Struct pointers
	var ms *myStruct
	// ms = &myStruct{id: 2, name: "Product Guyyy"}
	// fmt.Println(ms)
	ms = new(myStruct)
	(*ms).id = 2
	(*ms).name = "Similoluwa"
	fmt.Printf("%v, %T\n", *ms, *ms)

	c, d := 10, 25
	fmt.Println(c, d)
	swapInt(&c, &d)
	fmt.Println(c, d)

	// Zero-value of an uninitialized pointer is nil

}

/**
* Swap Integers using Pointers
 */
func swapInt(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}
