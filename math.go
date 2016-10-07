package main

import "fmt"

func f(n int) int {
	if n == 0 || n == 1 {
		return 1
	}	
	return n * f(n-1)
}

func main() {
	fmt.Println(f(10))
}
