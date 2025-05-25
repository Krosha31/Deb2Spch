package main

import (
	"html/template"
	"net/http"
	"os"

	"Deb2Spch/internal/auth"
	"Deb2Spch/internal/database"
	"Deb2Spch/internal/upload"
	"github.com/joho/godotenv"
)

var cwd string 

func mainPageHandler(w http.ResponseWriter, r *http.Request) { 
	template.Must(template.ParseFiles(cwd + "/web/html/main.html")).Execute(w, nil)
}

func main() {
	err := godotenv.Load()
    if err != nil {
        return 
    }
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return
	}
	auth.JwtSecret = []byte(secret)
	os.MkdirAll("uploads", os.ModePerm)
	database.Db = database.Database{}
	database.Db.Connect()
	defer database.Db.Disconnect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	cwd, _ = os.Getwd()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(cwd + "/web/css"))
    mux.Handle("/css/", http.StripPrefix("/css", fs))

	fs = http.FileServer(http.Dir(cwd + "/addons"))
    mux.Handle("/addons/", http.StripPrefix("/addons/", fs))

	fs = http.FileServer(http.Dir(cwd + "/web/scripts"))
    mux.Handle("/scripts/", http.StripPrefix("/scripts/", fs))

	fs = http.FileServer(http.Dir(cwd + "/web/common"))
    mux.Handle("/common/", http.StripPrefix("/common/", fs))


	mux.HandleFunc("/", mainPageHandler)
	mux.HandleFunc("/loginpage/", auth.LoginPageHandler)
	mux.HandleFunc("/login/", auth.LoginHandler)
	mux.HandleFunc("/registerpage/", auth.RegisterPageHandler)
	mux.HandleFunc("/register/", auth.RegisterHandler)
	mux.HandleFunc("/upload/", upload.UploadFileHandler)
	mux.HandleFunc("/split/", upload.SplitHandler)
	mux.HandleFunc("/refresh", auth.RefreshHandler)
	http.ListenAndServe(":" + port, mux)
}