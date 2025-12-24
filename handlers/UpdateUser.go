package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/models"
)

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)
	log.Printf("Updating profile for user ID: %d", uid)

	// Use a temporary struct to handle the date string
	var payload struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		DOB      *string `json:"dob"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Decode error: %v", err)
		jsonRes(w, map[string]string{"message": "Invalid input"}, 400)
		return
	}

	log.Printf("Received data: username=%s, email=%s, dob=%v", payload.Username, payload.Email, payload.DOB)

	// Parse the DOB if provided
	var dobTime *time.Time
	if payload.DOB != nil && *payload.DOB != "" {
		parsed, err := time.Parse("2006-01-02", *payload.DOB)
		if err != nil {
			log.Printf("Date parse error: %v", err)
			jsonRes(w, map[string]string{"message": "Invalid date format"}, 400)
			return
		}
		dobTime = &parsed
	}

	_, err := models.DB.Exec(`
		UPDATE users SET username=$1, email=$2, dob=$3
		WHERE id=$4
	`, payload.Username, payload.Email, dobTime, uid)

	if err != nil {
		log.Printf("Database update error: %v", err)
		jsonRes(w, map[string]string{"message": "Update failed"}, 500)
		return
	}

	log.Println("Profile updated successfully")
	jsonRes(w, map[string]string{"message": "Profile updated"}, 200)
}
