package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey There</h1>
					<p>Go is fast!</p>`)
	fmt.Fprintf(w, "<p>Formated string: %s</p>", "<strong>This is a string</strong>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}