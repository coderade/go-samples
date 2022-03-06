package main

import "fmt"

func main() {

	// Different ways of creating a map
	colors1 := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
	}
	fmt.Println(colors1)

	//var colors2 map[string]string
	//colors2["blue"] = "#3389FF"
	//fmt.Println(colors2)

	//colors3 := make(map[string]string)
	//colors3["white"] = "#FFFFFF"
	//fmt.Println(colors3)

	delete(colors1, "red")
	fmt.Println(colors1)
}
