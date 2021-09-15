package main

import (
	"fmt"
	"net/http"
)

func main() {
	serve("<h1>Hello world, this is my first webserver in Golang</>")
}

func serve(content string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //callback
		w.Write([]byte(content))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Server running on PORT 8080.")
	}
}
