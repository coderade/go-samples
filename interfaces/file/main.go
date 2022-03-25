package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args
	file := args[1]
	dir, _ := filepath.Abs("./")
	filePath := dir + "/" + file
	data, err := os.ReadFile(filePath)
	check(err)
	fmt.Println(string(data))
}
