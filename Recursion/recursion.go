package main

import "fmt"

// factorial
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// fibonaci
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
		// fib(7) = fib(6) + fib(5) = 13
		// fib(6) = fib(5) + fib(4) = 8
		// fib(5) = fin(4) + fib(3) = 5
		// fib(4) = fib(3) + fib(2) = 3
		// fib(3) = fib(2) + fib(1) = 2
		// fib(2) = fib(1) + fib(0) = 1
	}
	fmt.Println(fib(7))
}
