package main

import (
	"log"
	"net/http"
	"todo/handlers"
	"todo/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()

	router := mux.NewRouter()

	// Simple logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("INCOMING: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Authentication routes
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Task routes (protected)
	router.HandleFunc("/api/tasks", handlers.AuthMiddleware(handlers.CreateTask)).Methods("POST")
	router.HandleFunc("/api/tasks", handlers.AuthMiddleware(handlers.GetTasks)).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", handlers.AuthMiddleware(handlers.UpdateTask)).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handlers.AuthMiddleware(handlers.DeleteTask)).Methods("DELETE")

	// Profile routes (protected)
	router.HandleFunc("/api/user/profile", handlers.AuthMiddleware(handlers.GetUserProfile)).Methods("GET")
	router.HandleFunc("/api/user/profile", handlers.AuthMiddleware(handlers.UpdateUserProfile)).Methods("PUT")

	// Serve static frontend files from the current folder
	fs := http.FileServer(http.Dir("./"))
	router.PathPrefix("/").Handler(fs)

	// --- KEEPING YOUR OLD DEBUG LOGS ---
	log.Println("Routes registered:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		log.Println(t)
		return nil
	})
	

	log.Println("🚀 Server running at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
