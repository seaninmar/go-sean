// https://go.dev/tour/methods/23

// Exercise: rot13Reader

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream
// in some way.

// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed
// data) and returns a *gzip.Reader that also implements io.Reader (a stream of the
// decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
// modifying the stream by applying the rot13 substitution cipher to all alphabetical
// characters.

// The rot13Reader type is provided for you. Make it an io.Reader by implementing its
// Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
