package main

import "fmt"

func main() {

	// default value is zero (0)
	var a [5]int
	fmt.Println("emp:", a)

	// set value 100 to index 4 in arrays
	a[4] = 100
	fmt.Println("set:", a)
	// getting numbers in an array
	fmt.Println("get:", a[4])
	// getting length
	fmt.Println("len:", len(a))

	// initialize array contents
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2nd:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2nd:", twoD)

	var threeD [3][4]int
	threeD = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("3nd:", threeD)
}
