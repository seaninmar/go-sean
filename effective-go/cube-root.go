package main

import (
	"fmt"
	"math"
)

// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
	z := x / 3 // Arbitrary initial value
	for i := 0; i < 1e6; i++ {
		prevz := z
		z -= (z*z*z - x) / (3 * z * z)
		if math.Abs(z-prevz) < 1e-6 {
			fmt.Print(i, ":")
			return z
		}
	}
	// A million iterations has not converged; something is wrong.
	panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

func main() {
	fmt.Println("CubeRoot(3) =", CubeRoot(3))
	fmt.Println("CubeRoot(-1) =", CubeRoot(-1))
	fmt.Println("CubeRoot(-10) =", CubeRoot(-10))
	fmt.Println("CubeRoot(0) =", CubeRoot(0)) // will panic because float is weird
}
