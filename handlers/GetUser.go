package handlers

import (
	"log"
	"net/http"
	"todo/models"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)
	log.Printf("Fetching profile for user ID: %d", uid)

	var user models.User
	err := models.DB.QueryRow(`
		SELECT id, username, email, dob, created_at
		FROM users WHERE id=$1
	`, uid).Scan(&user.ID, &user.Username, &user.Email, &user.DOB, &user.CreatedAt)

	if err != nil {
		log.Printf("Database error: %v", err)
		jsonRes(w, map[string]string{"message": "User not found"}, 404)
		return
	}

	log.Printf("User found: %+v", user)
	jsonRes(w, user, 200)
}
