package main

import "fmt"

var colors = [5]string{"red", "green", "blue"}
var systems = []string{"linux", "macos", "windows"}

func main() {
	fmt.Println(NumberOfColors(), colors)
	fmt.Println(NumberOfSystems(), systems)
}

// NumberOfColors ..
func NumberOfColors() int {
	return len(colors)
}

// NumberOfSystems ..
func NumberOfSystems() int {
	return len(systems)
}
