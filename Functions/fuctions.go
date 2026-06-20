package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func times(x, b int) int {
	return x * b
}

func divide(x, y int) int {
	return x / y
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	res = times(2, 5)
	fmt.Println("2 x 5 =", res)

	var divided = divide(20, 4)
	fmt.Println("20 / 4 =", divided)
}
