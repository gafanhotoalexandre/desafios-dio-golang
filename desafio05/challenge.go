package main

import "fmt"

func main() {
	// operations
	resSum := sum(5, 4, 3, 2, 1)
	resSubtract := subtract(100, 11, 11)
	resMultiply := multiply(1, 5, 2)
	resDivide := divide(100, 2, 2)

	fmt.Println(resSum)
	fmt.Println(resSubtract)
	fmt.Println(resMultiply)
	fmt.Println(resDivide)
}

// functions
func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func subtract(i ...int) int {
	if len(i) == 0 {
		return 0
	}

	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func multiply(i ...int) int {
	if len(i) == 0 {
		return 0
	}

	total := i[0]
	for _, v := range i[1:] {
		total *= v
	}
	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}

	total := i[0]
	for _, v := range i[1:] {
		total /= v
	}
	return total
}
