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
	var managerName *string

	// Get user data and their manager's name with LEFT JOIN
	err := models.DB.QueryRow(`
		SELECT u.id, u.username, u.email, u.dob, u.created_at, u.manager_id, m.username
		FROM users u
		LEFT JOIN users m ON u.manager_id = m.id
		WHERE u.id=$1
	`, uid).Scan(&user.ID, &user.Username, &user.Email, &user.DOB, &user.CreatedAt, &user.ManagerID, &managerName)

	if err != nil {
		log.Printf("Database error: %v", err)
		jsonRes(w, map[string]string{"message": "User not found"}, 404)
		return
	}

	log.Printf("User found: %+v", user)

	// Create response with manager name
	response := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"dob":        user.DOB,
		"created_at": user.CreatedAt,
		"manager_id": user.ManagerID,
	}

	if managerName != nil {
		response["manager_name"] = *managerName
	}

	jsonRes(w, response, 200)
}