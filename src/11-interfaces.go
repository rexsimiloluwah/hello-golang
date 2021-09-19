package main

import (
	"fmt"
	"math"
)

type Writer interface {
	Write([]byte) (int, error)
}
type tank interface {
	Area() float64
	Volume() float64
}

// Single method interface
type Incrementer interface {
	Increment() int
}

func main() {

	var cwriter Writer = ConsoleWriter{} // ConsoleWriter struct implements the Writer interface 
	cwriter.Write([]byte("This is the movement"))

	var rt, ct tank

	// int implementation
	myInt := IntCounter(0)
	var inc Incrementer = &myInt;

	for i:= 0; i<10; i++{
		fmt.Println(inc.Increment())
	}

	
	r := RectangularTank{
		1.5,
		2.3,
		5.0,
	}
	c := CylindricalTank{
		5.4,
		1.2,
	}

	// The interfaces are implicitly implemented
	rt = r // RectangularTank struct implements the tank interface here
	fmt.Printf("The area of the rectangular tank is: - %.2f cm2\n", rt.Area())
	fmt.Printf("The volume of the rectangular tank is: - %.2f cm3\n", rt.Volume())

	ct = c // CylindricalTank struct implements the tank interface here
	fmt.Printf("The area of the cylindrical tank is :- %.2f cm2\n", ct.Area())
	fmt.Printf("The volume of the cylindrical tank is :- %.2f cm3\n", ct.Volume())

	// Empty interfaces 
	var val interface{} = 1.2
	typeof(val)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(content []byte) (int,error){
	n,err := fmt.Println(content);
	return n,err;
}

type RectangularTank struct {
	l float64
	b float64
	h float64
}

type CylindricalTank struct {
	r float64
	h float64
}

// Using integers to implement interfaces
type IntCounter int 

func (r RectangularTank) Area() float64 {
	return 2 * (r.l*r.b + r.l*r.h + r.b*r.h)
}

func (r RectangularTank) Volume() float64 {
	return r.l * r.b * r.h
}

func (c CylindricalTank) Area() float64 {
	return 2 * (math.Pi*math.Pow(c.r, 2) + (math.Pi * c.r * c.h))
}

func (c CylindricalTank) Volume() float64 {
	return math.Pi * math.Pow(c.r, 2) * c.h
}

func (ic *IntCounter) Increment() int{
	*ic++ 
	return int(*ic)
}

// Type switches using empty interfaces
func typeof(v interface{}) {
	switch v.(type){
	case int:
		fmt.Println("This is an Integer")
	case float64:
		fmt.Println("This is a Float")
	case bool:
		fmt.Println("This is a Boolean")
	case string:
		fmt.Println("This is a String")
	default:
		fmt.Println("I don't know about thisss")
	}
}