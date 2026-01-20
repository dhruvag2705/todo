package main

import (
	"log"	//logs for server activities
	"net/http"	//HTTP server and request handling
	"todo/handlers"	//handler functions for API endpoints
	"todo/models" //database models and connection

	"github.com/gorilla/mux" //router for handling HTTP routes
)

func main() {
	models.ConnectDB() // Connect to the database

	router := mux.NewRouter() // Create a new router

	// Simple logging middleware for incoming requests
	router.Use(func(next http.Handler) http.Handler {  //
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //handler function 
			log.Printf("INCOMING: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Authentication routes
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/users", handlers.GetAllUsers).Methods("GET")		// New route to get all users

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
