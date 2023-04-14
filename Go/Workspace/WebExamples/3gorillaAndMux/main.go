package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			//get the book
			title := vars["title"]
			//navigate to the page
			page := vars["page"]
			fmt.Fprintf(w,
				"You have requested the book with the name %s on the page %s",
				title, page)
		})

	http.ListenAndServe(":80", r)
}
