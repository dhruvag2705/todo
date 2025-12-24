package handlers

import (
	"encoding/json"
	"net/http"

	"todo/models"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)
	params := mux.Vars(r)
	id := params["id"]

	var body struct {
		Completed bool `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonRes(w, "invalid body", http.StatusBadRequest)
		return
	}

	var t models.Task
	err := models.DB.QueryRow(
		`UPDATE tasks 
		 SET completed=$1, completed_at=CASE WHEN $1 THEN NOW() ELSE NULL END
		 WHERE id=$2 AND user_id=$3
		 RETURNING id, title, due_date, completed, completed_at`,
		body.Completed, id, uid,
	).Scan(&t.ID, &t.Title, &t.DueDate, &t.Completed, &t.CompletedAt)

	if err != nil {
		jsonRes(w, "update failed", http.StatusInternalServerError)
		return
	}

	jsonRes(w, t, http.StatusOK)
}
