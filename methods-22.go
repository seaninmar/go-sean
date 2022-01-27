// https://go.dev/tour/methods/22

// Exercise: Readers

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
