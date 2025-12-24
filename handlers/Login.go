package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"todo/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret")

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("=== LOGIN ATTEMPT STARTED ===")

	var creds models.Credentials // Changed from models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println("ERROR: Decode failed:", err)
		jsonRes(w, map[string]string{"message": "Invalid input"}, 400)
		return
	}

	log.Println("Username received:", creds.Username)
	log.Println("Password received:", creds.Password)

	var id int
	var pass string
	err := models.DB.QueryRow("SELECT id, password FROM users WHERE username=$1",
		creds.Username).Scan(&id, &pass)

	if err != nil {
		log.Println("ERROR: User not found in database:", err)
		jsonRes(w, map[string]string{"message": "User not found"}, 404)
		return
	}

	log.Println("SUCCESS: User found, ID:", id)
	log.Println("Stored hash length:", len(pass))

	// DETAILED COMPARISON LOG
	if len(pass) > 30 {
		log.Printf("Hash (first 30 chars): %s...", pass[:30])
	} else {
		log.Printf("Hash: %s", pass)
	}
	log.Printf("Password entered: '%s'", creds.Password)

	// PASSWORD CHECK - THIS IS THE KEY PART
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(creds.Password))
	if err != nil {
		log.Println("ERROR: Password check FAILED:", err)
		jsonRes(w, map[string]string{"message": "Wrong password"}, 401)
		return
	}

	log.Println("SUCCESS: Password matched!")

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}).SignedString(jwtKey)

	log.Println("SUCCESS: Token created")
	log.Println("=== LOGIN COMPLETED ===")
	jsonRes(w, map[string]string{"token": token}, 200)
}
