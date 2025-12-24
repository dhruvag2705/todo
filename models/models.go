package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	DOB       *time.Time `json:"dob,omitempty"`
	Password  string    `json:"-"` // password protection in JSON response
	CreatedAt time.Time  `json:"created_at"`
}
 
type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID          int        `json:"id"`
	UserID      int        `json:"-"`
	Title       string     `json:"title"`
	DueDate     time.Time  `json:"dueDate"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completedAt"`
}
