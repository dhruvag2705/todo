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

	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println("ERROR: Decode failed:", err)
		jsonRes(w, map[string]string{"message": "Invalid input"}, 400)
		return
	}

	// Validate required fields
	if creds.Email == "" {
		jsonRes(w, map[string]string{"message": "Email is required"}, 400)
		return
	}

	// Look up manager ID if manager username is provided
	var managerID *int
	if creds.ManagerUsername != "" {
		var mgrID int
		err := models.DB.QueryRow("SELECT id FROM users WHERE username=$1", creds.ManagerUsername).Scan(&mgrID)
		if err != nil {
			log.Println("Manager not found:", err)
			jsonRes(w, map[string]string{"message": "Manager not found"}, 400)
			return
		}
		managerID = &mgrID
	}

	// Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 10)

	// Insert new user with manager_id
	if _, err := models.DB.Exec("INSERT INTO users(username, email, password, manager_id) VALUES ($1, $2, $3, $4)",
		creds.Username, creds.Email, string(hash), managerID); err != nil {

		log.Println("❌ Error inserting user:", err)
		jsonRes(w, map[string]string{"message": "Error creating user"}, 400)
		return
	}

	log.Println("✅ User created successfully")
	jsonRes(w, map[string]string{"message": "User created"}, 200)
}