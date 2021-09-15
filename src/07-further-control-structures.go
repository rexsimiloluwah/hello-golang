package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Further control structures :- defer, panic, recovery
func main() {
	// Learning Defer
	// fetchFromFile("file.txt")
	// fetchFromUrl("http://www.google.com/robots.txt")

	// Learning Panic
	panicFunc()
}

func fetchFromFile(filepath string) {
	data := make([]byte, 200) // Slice to store the data read from the file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close() // This defers the file closing operation to execute after the function's context
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Read %d bytes\nContent:- %s\n", count, data[:count])
}

func fetchFromUrl(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", content)
}

func panicFunc() {
	// Playground for understanding panic and exception handling
	// Panic is basically used for throwing errors during exception handling
	// a, b := 1, 0
	// c := a / b // Exception due to division error
	// fmt.Println(c)

	// fmt.Println("start")
	// panic("Something bad happened")
	// fmt.Println("end")

}
