package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

// Sum ...
func Sum(a,b,c,d,e int) int {
	s := (a+b+c+d+e)
	return s
}