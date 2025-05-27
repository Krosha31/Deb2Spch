package main

import (
	"net/http"
	"os"

	"Deb2Spch/internal/auth"
	"Deb2Spch/internal/database"
	"Deb2Spch/internal/upload"
	"Deb2Spch/internal/pages"
	"github.com/joho/godotenv"
)

var cwd string 

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
	database.Db = database.Database{}
	database.Db.Connect()
	defer database.Db.Disconnect()
	port := "3000"
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

	// mux.Handle("/deb2spch/", http.StripPrefix("/deb2spch/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/", pages.MainPageHandler)
	mux.HandleFunc("/profile/", pages.ProfilePageHandler)
	mux.HandleFunc("/loginpage/", pages.LoginPageHandler)
	mux.HandleFunc("/registerpage/", pages.RegisterPageHandler)
	mux.HandleFunc("/subscription/", pages.SubscriptionPageHandler)


	mux.HandleFunc("/login/", auth.LoginHandler)
	mux.HandleFunc("/register/", auth.RegisterHandler)
	mux.HandleFunc("/upload/", upload.UploadFileHandler)
	mux.HandleFunc("/split/", upload.SplitHandler)
	mux.HandleFunc("/refresh/", auth.RefreshHandler)

	http.ListenAndServe(":" + port, mux)
}