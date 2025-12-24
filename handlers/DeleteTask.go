package handlers

import (
	"net/http"
	"todo/models"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userId").(int)
	id := mux.Vars(r)["id"]

	if _, err := models.DB.Exec(
		"DELETE FROM tasks WHERE id=$1 AND user_id=$2",
		id, uid); err != nil {

		jsonRes(w, "err", 500)
		return
	}

	jsonRes(w, "Task deleted", 200)
}
