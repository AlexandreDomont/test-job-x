package math
package main

import "fmt"

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

func main() {
	fmt.Println(Add(1, 1))
}