package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog"})
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	http.ListenAndServe(":8000", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
