// methods/20
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.
	i := 0
	err := 0.00000000001
	for ; i < 10; i++ {
		new_z := z - ((z*z - x) / (2*z))
		diff := new_z - z
		if diff < 0 {
			diff *= -1
		}
		if diff < err {
			return new_z, nil
		} else {
			z = new_z
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
