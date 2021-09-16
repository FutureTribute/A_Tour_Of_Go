// flowcontrol/8
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.
	i := 0
	err := 0.00000000001
	for ; i < 10; i++ {
		fmt.Println("i", i)
		new_z := z - ((z*z - x) / (2*z))
		diff := new_z - z
		if diff < 0 {
			diff *= -1
		}
		fmt.Println("Diff", diff)
		if diff < err {
			return new_z
		} else {
			z = new_z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
