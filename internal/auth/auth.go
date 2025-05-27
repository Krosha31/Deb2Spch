package auth

import (
	"Deb2Spch/internal/common"
	"Deb2Spch/internal/database"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type loginPassword struct {
	Login    string
	Password string
}

type Claims struct {
	Login string
	jwt.RegisteredClaims
}

var JwtSecret []byte

func getTokens(login string) (string, string){
	accessClaims := &Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
		  ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	  }
	  accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(JwtSecret)
	
	  refreshClaims := &Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
		  ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	  }
	  refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(JwtSecret)
	  return accessToken, refreshToken
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
	  http.Error(w, "Нет refresh токена", http.StatusUnauthorized)
	  return
	}
  
	refreshTokenStr := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshTokenStr, claims, func(token *jwt.Token) (interface{}, error) {
	  return JwtSecret, nil
	})
	if err != nil || !token.Valid {
	  http.Error(w, "Неверный refresh токен", http.StatusUnauthorized)
	  return
	}
  
	newAccessClaims := &Claims{
	  Login: claims.Login,
	  RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	  },
	}
	newAccessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessClaims).SignedString(JwtSecret)
  
	json.NewEncoder(w).Encode(map[string]string{"token": newAccessToken})
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
	fmt.Println(jsonBody.Login)
	user, err := database.Db.GetUserByLogin(jsonBody.Login)
	if err != nil {
		http.Error(w, "Error getting access to database", http.StatusInternalServerError)
		return
	}
	if user == (common.User{}) {
		http.Error(w, "User doesn't exist", http.StatusNotFound)
		return 
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(jsonBody.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusNotFound)
		return
	}

	accessToken, refreshToken := getTokens(jsonBody.Login)

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false, // true на проде с HTTPS
		SameSite: http.SameSiteStrictMode,
		Path:     "/deb2spch/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	json.NewEncoder(w).Encode(map[string]string{"token": accessToken})
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
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	err = database.Db.AddUser(jsonBody.Login, string(hashedPassword))
	if err != nil{
		fmt.Println("database", err)
		http.Error(w, "Error getting access to database", http.StatusInternalServerError)
		return
	}

	accessToken, refreshToken := getTokens(jsonBody.Login)

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false, // true на проде с HTTPS
		SameSite: http.SameSiteStrictMode,
		Path:     "/deb2spch/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	json.NewEncoder(w).Encode(map[string]string{"token": accessToken})
}
