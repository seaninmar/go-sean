// https://www.golang-book.com/books/intro/10

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(ch chan string, n int, o int) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprint(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	if n == o-1 {
		close(ch)
	}
}

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go f(ch, i, 10)
	}
	for msg := range ch {
		fmt.Print(msg, " ")
	}
	fmt.Println()
}
