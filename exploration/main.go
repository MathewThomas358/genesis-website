package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	//So that i can quickly change what im rendering

	page := whole_page(10, 10)

	//Serving in /about
	http.Handle("/about", templ.Handler(page))
	//Serving css in /static
	http.Handle("/static/", http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe("localhost:3000", nil)

}
