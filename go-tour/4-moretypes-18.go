// https://go.dev/tour/moretypes/18

// Exercise: Slices

// Implement Pic. It should return a slice of length dy, each element of which is a
// slice of dx 8-bit unsigned integers. When you run the program, it will display your
// picture, interpreting the integers as grayscale (well, bluescale) values.

// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and
// x^y.

// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)

// (Use uint8(intValue) to convert between types.)

package main

import "golang.org/x/tour/pic"

type PicFunc func(int, int) [][]uint8
type ImageFunc func(int, int) uint8

func PicFactory(f ImageFunc) PicFunc {
	picFunc := func(dx, dy int) [][]uint8 {
		// allocate y
		y := make([][]uint8, dy)
		for i := range y {
			y[i] = make([]uint8, dx)
			// draw the picture
			for j := range y[i] {
				y[i][j] = f(i, j)
			}
		}
		return y
	}
	return picFunc
}

func f1(x, y int) uint8 {
	return uint8((x + y) / 2)
}
func f2(x, y int) uint8 {
	return uint8(x * y)
}
func f3(x, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	// pick 1: f1, f2, f3 (it would be nice to define f1/2/3 as a lambda inline inside
	// PicFactory(...), but not in Go apparently)
	pic.Show(PicFactory(f1))
}
