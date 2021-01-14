package main

import (
	"fmt"
	"net/http"
)

func main() {
	//static file handler
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)

	http.HandleFunc("/mynameis", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Forbidden!")
		} else if r.FormValue("name") == "" {
			fmt.Fprintf(w, "No name was informed!")
		} else {
			fmt.Fprintf(w, "Your name  is %s!", r.FormValue("name"))
		}
	})

	http.ListenAndServe(":80", nil)
}
