package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
