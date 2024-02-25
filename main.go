package main

import "fmt"

func main() {

	var s string
	var t string

	s = "HI!"
	t = "HI!"

	if s == t {
		fmt.Println("Same")
	}

	if s != t {
		fmt.Println("Diff")
	}
}
