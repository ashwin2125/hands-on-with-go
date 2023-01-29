package main

import "fmt"

func main() {
	Greet("Alice")
	Greet("Bob")
}

// Greet ...
func Greet (name string){
	fmt.Println("Hello,",name)
}