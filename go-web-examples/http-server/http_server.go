package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!\n")

		// This verify if there is a GET parameter "msg" and returns it for the user.
		if r.URL.Query().Get("msg") != "" {
			fmt.Fprintf(w, "You sent this message: "+r.URL.Query().Get("msg"))
		}
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}
