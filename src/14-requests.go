package main 

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Map map[string]interface{}

func main(){
	// Testing GET Request 
	//fetchData()

	// Testing the POST Request
	testData:= Map{
		"name": "Similoluwa Okunowo",
		"email": "rexsimiloluwa@gmail.com",
		"age": 20,
		"school": "Obafemi Awolowo University, Ile-ife",
		"timestamp": "2021-09-19 01:42:00 PM",
	}

	//postData(testData)

	
}

/* Get request */
func fetchData(){
	resp, err:= http.Get("https://jsonplaceholder.typicode.com/posts")
	if(err != nil){
		log.Fatalln(err)
		return
	}

	// Close the response body at the end of the fetchData() execution stack
	defer resp.Body.Close()
	body,err:= ioutil.ReadAll(resp.Body)
	if(err != nil){
		log.Fatalln(err)
		return
	}

	fmt.Println(string(body))
}

/* POST Request */
func postData(data Map){
	// Marshal is used to perform JSON encoding of the data
	d, _ := json.Marshal(data)
	// Convert to io R
	postBody:= bytes.NewBuffer(d)
	// Making the request using http's Post method 
	// Use a http client for more control over the headers https://pkg.go.dev/net/http
	resp, err:= http.Post(
		"https://postman-echo.com/post",
		"application/json",
		postBody,
	)
	defer resp.Body.Close()
	if err != nil{
		log.Fatalln(err)
		return
	}
	// Read the response body 
	body,err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
		return
	}

	fmt.Println(string(body))
}

/* PUT/PATCH Request */

/* DELETE Request */