// hello.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "tui la phung ne %s\n", r.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
