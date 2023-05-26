package main

import "fmt"

func main() {
	fmt.Println("welcome")
	a := 5
	b := 6
	fmt.Println(a * b)
	fmt.Println(sub(a, b))
}

func sub(a int, b int) int {
	return b - a
}

// just a comment
