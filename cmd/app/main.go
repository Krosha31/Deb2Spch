package main

import (
	"html/template"
	"net/http"
	"os"

	"Deb2Spch/internal/auth"
	"Deb2Spch/internal/database"
)



var mainPageHtml = template.Must(template.ParseFiles("web/html/main.html"))

func mainPageHandler(w http.ResponseWriter, r *http.Request) { 
	mainPageHtml.Execute(w, nil)
}

func main() {
	db := database.Database{}
	db.Connect()
	defer db.Disconnect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/css"))
    mux.Handle("/css/", http.StripPrefix("/css", fs))

	fs = http.FileServer(http.Dir("addons"))
    mux.Handle("/addons/", http.StripPrefix("/addons/", fs))

	fs = http.FileServer(http.Dir("./web/scripts"))
    mux.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	fs = http.FileServer(http.Dir("./web/common"))
    mux.Handle("/common/", http.StripPrefix("/common/", fs))


	mux.HandleFunc("/", mainPageHandler)
	mux.HandleFunc("/loginpage/", auth.LoginPageHandler)
	mux.HandleFunc("/login/", auth.LoginHandler)
	mux.HandleFunc("/register/", auth.LoginPageHandler)
	http.ListenAndServe(":" + port, mux)
}