package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)

	var body struct {
		Title   string `json:"title"`
		DueDate string `json:"dueDate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("ERROR: JSON decode failed:", err)
		jsonRes(w, "bad json", 400)
		return
	}

	log.Println("Received dueDate:", body.DueDate)

	// Parse ISO 8601 format with timezone
	parsed, err := time.Parse(time.RFC3339, body.DueDate)
	if err != nil {
		log.Println("ERROR: Date parse failed:", err)
		jsonRes(w, "Invalid date format", 400)
		return
	}

	log.Println("Parsed datetime:", parsed)

	t := models.Task{
		Title:   body.Title,
		DueDate: parsed,
	}

	err = models.DB.QueryRow(
		`INSERT INTO tasks (user_id, title, due_date)
		 VALUES ($1,$2,$3) RETURNING id, completed`,
		uid, t.Title, t.DueDate,
	).Scan(&t.ID, &t.Completed)

	if err != nil {
		log.Println("ERROR: Database insert failed:", err)
		jsonRes(w, err.Error(), 500)
		return
	}

	log.Println("Task created successfully:", t.ID)
	jsonRes(w, t, 200)
}
