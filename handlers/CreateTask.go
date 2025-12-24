package handlers

import (
	"encoding/json"
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
		jsonRes(w, "bad json", 400)
		return
	}

	t := models.Task{
		Title: body.Title,
	}

	parsed, err := time.Parse("2006-01-02", body.DueDate)
	if err == nil {
		t.DueDate = parsed
	}

	err = models.DB.QueryRow(
		`INSERT INTO tasks (user_id, title, due_date)
		 VALUES ($1,$2,$3) RETURNING id, completed`,
		uid, t.Title, t.DueDate,
	).Scan(&t.ID, &t.Completed)

	if err != nil {
		jsonRes(w, err.Error(), 500)
		return
	}

	jsonRes(w, t, 200)
}
