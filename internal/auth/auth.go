package auth

import (
	"fmt"
	"net/http"
	"html/template"
	"io"
	// "golang.org/x/crypto/bcrypt"
	"time"
)

type user struct {
	id 					int
	login 				string 
	password_hash 		string
	subscribtion_id 	int
	registration_date 	time.Time	
}


func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		defer req.Body.Close()

		fmt.Fprintf(w, "Полученные данные: %s\n", body)
	} else {
		template.Must(template.ParseFiles("web/html/auth.html")).Execute(w, nil)
	}
	
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("rth")
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	fmt.Println(w, "Полученные данные: %s\n", body)
}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	
}