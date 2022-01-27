// https://go.dev/tour/methods/20

// Exercise: Errors

// Copy your Sqrt function from the earlier exercise and modify it to return an error
// value.

// Sqrt should return a non-nil error value when given a negative number, as it doesn't
// support complex numbers.

// Create a new type

// type ErrNegativeSqrt float64

// and make it an error by giving it a

// func (e ErrNegativeSqrt) Error() string

// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number:
// -2".

// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an
// infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative
// number.

package main

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf(
		"cannot Sqrt negative number: %v", 	
		float64(e),
	)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegSqrt(x)
	}
	var z1 float64 = 1
	var z2 = x
	var i int
	for math.Abs(z2-z1) > 1e-15 {
		z1 = z2
		z2 = newGuess(z1, x)
		i += 1
	}
	return z2, nil
}

func newGuess(z, x float64) float64 {
	return z - ((z*z - x) / (2 * z))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
