package handlers

import (
	"net/http"

	"todo/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)

	rows, _ := models.DB.Query("SELECT id, title, due_date, completed, completed_at FROM tasks WHERE user_id=$1 ORDER BY id DESC", uid)
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.DueDate, &t.Completed, &t.CompletedAt)
		tasks = append(tasks, t)
	}

	jsonRes(w, tasks, 200)
}
