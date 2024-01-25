package main

import "fmt"

func main() {
	var a, b = 3, 6
	fmt.Printf("Binary exponentiation: %d in power %d equals %d\n", a, b, binaryExp(a, b))
}

// Returns number^power value
// calculated by bitwise shift of power
// with O(n) = log n
func binaryExp(number, power int) int {
	if power == 0 {
		return 1
	}

	res := 1
	for power > 0 {
		if power&1 == 1 {
			res *= number
		}
		number *= number
		power >>= 1
	}

	return res
}
