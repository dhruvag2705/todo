package handlers

import (
	"log"
	"net/http"
	"todo/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := models.DB.Query("SELECT id, username FROM users ORDER BY username")
	if err != nil {
		log.Println("Error fetching users:", err)
		jsonRes(w, map[string]string{"message": "Error fetching users"}, 500)
		return
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var username string
		rows.Scan(&id, &username)
		users = append(users, map[string]interface{}{
			"id":       id,
			"username": username,
		})
	}

	jsonRes(w, users, 200)
}