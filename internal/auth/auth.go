package auth

import (
	"Deb2Spch/internal/common"
	"Deb2Spch/internal/database"
	// . "Deb2Spch/internal/common"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	// "golang.org/x/crypto/bcrypt"
	"encoding/json"
)

type loginPassword struct {
	Login    string
	Password string
}

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/auth.html")).Execute(w, nil)
}

func RegisterPageHandler(w http.ResponseWriter, req *http.Request) {
	cwd, _ := os.Getwd()
	template.Must(template.ParseFiles(cwd + "/web/html/register.html")).Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var jsonBody loginPassword
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	user, err := database.Db.GetUserByLogin(jsonBody.Login)
	if err != nil {
		http.Error(w, "Error getting access to database", http.StatusInternalServerError)
		return
	}
	if user == (common.User{}) {
		http.Error(w, "User doesn't exist", http.StatusNotFound)
		return 
	}
	w.WriteHeader(http.StatusOK)
}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var jsonBody loginPassword
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
	}

	user, err := database.Db.GetUserByLogin(jsonBody.Login)
	if err != nil {
		http.Error(w, "Error getting access to database", http.StatusInternalServerError)
	}
	if user != (common.User{}) {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	
	err = database.Db.AddUser(jsonBody.Login, jsonBody.Password)
	if err != nil{
		fmt.Println("database", err)
		http.Error(w, "Error getting access to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
