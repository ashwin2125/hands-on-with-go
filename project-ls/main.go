package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	name := (strings.Join(listFiles("testdata"), " "))
	fmt.Println(name)
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
