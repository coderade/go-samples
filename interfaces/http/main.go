package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Creating a byte slice and printing the response.body
	//bs := make([]byte, 99999)
	//response.Body.Read(bs)
	//fmt.Println(string(bs))

	//Using the io.Copy from os.Stdout to print the response.body
	io.Copy(os.Stdout, response.Body)

}
