package models

import "time"		// For time handling

type User struct { 				// For users
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	DOB       *time.Time `json:"dob,omitempty"`
	Password  string     `json:"-"` // password protection in JSON response	
	ManagerID *int       `json:"manager_id,omitempty"`		// For hierarchical user relationships
	CreatedAt time.Time  `json:"created_at"`
}

type Credentials struct { 		// For login/signup
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ManagerUsername string `json:"manager_username"`	// For assigning a manager during signup
}

type Task struct {  			// For tasks
	ID          int        `json:"id"`
	UserID      int        `json:"-"`
	Title       string     `json:"title"`
	DueDate     time.Time  `json:"dueDate"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completedAt"`
}
