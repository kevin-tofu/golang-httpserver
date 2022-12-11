package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":8008",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
		}),
	)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

}
