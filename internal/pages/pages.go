package pages

import (
	"html/template"
	"net/http"
	"os"
)

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/auth.html")).Execute(w, nil)
}

func RegisterPageHandler(w http.ResponseWriter, req *http.Request) {
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/register.html")).Execute(w, nil)
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) { 
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/main.html")).Execute(w, nil)
}

func ProfilePageHandler(w http.ResponseWriter, r *http.Request) { 
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/profile.html")).Execute(w, nil)
}

func SubscriptionPageHandler(w http.ResponseWriter, r *http.Request) { 
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/subscription.html")).Execute(w, nil)
}