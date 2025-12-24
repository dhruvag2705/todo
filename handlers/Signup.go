package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/models"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("Signup endpoint hit")

	var creds models.Credentials // Changed from models.User
	if json.NewDecoder(r.Body).Decode(&creds) != nil {
		jsonRes(w, "Invalid input", 400)
		return
	}

	// Validate that email is provided
	if creds.Email == "" {
		jsonRes(w, map[string]string{"message": "Email is required"}, 400)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 10)

	if _, err := models.DB.Exec("INSERT INTO users(username, email, password) VALUES ($1, $2, $3)",
		creds.Username, creds.Email, string(hash)); err != nil {

		log.Println("❌ Error inserting user:", err)
		jsonRes(w, map[string]string{"message": "Error creating user"}, 400)
		return
	}

	jsonRes(w, map[string]string{"message": "User created"}, 200)
}
