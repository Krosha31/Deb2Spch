package main

import (
	"net/http"
	"html/template"
	"os"
)

var mainPageHtml = template.Must(template.ParseFiles("html/main.html"))

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	mainPageHtml.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("css"))
    mux.Handle("/css/", http.StripPrefix("/css/", fs))

	fs = http.FileServer(http.Dir("addons"))
    mux.Handle("/addons/", http.StripPrefix("/addons/", fs))

	mux.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":" + port, mux)
}