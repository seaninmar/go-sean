package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Counter int

func (ctr *Counter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	*ctr++
	fmt.Fprintf(response, "counter = %d\n", ctr)
}

// A channel that sends a notification on each visit
type Channel chan *http.Request

func (ch Channel) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ch <- request
	fmt.Fprint(response, "notification sent")
}

// An HTTP endpoint to display the calling arguments
func ArgServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, os.Args)
}

func main() {
	http.Handle("/args", http.HandlerFunc(ArgServer))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
