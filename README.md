
# 📝 Todo Web Application (Go + HTML)

A **full-stack Todo Management Web Application** built using **Go (Golang)** for the backend and **HTML** for the frontend.
It allows users to **sign up, log in, and manage their tasks securely** with full CRUD functionality.

---

## 🚀 Features

* User **Signup & Login** (Authentication)
* **Create, Read, Update, Delete (CRUD)** Todo tasks
* User profile management
* Task progress tracking
* JSON-based API responses
* Middleware for request handling
* Clean modular project structure

---

## 🛠️ Technologies Used

* **Backend:** Go (Golang)
* **Frontend:** HTML
* **Database:** SQL (configured in `db.go`)
* **Architecture:** MVC-style folder structure
* **Authentication:** Custom login & signup handlers

---

## 📂 Project Structure

```
todo/
│── main.go
│── go.mod
│── go.sum
│
├── handlers/
│   ├── Signup.go
│   ├── Login.go
│   ├── CreateTask.go
│   ├── GetTasks.go
│   ├── UpdateTask.go
│   ├── DeleteTask.go
│   ├── GetUser.go
│   ├── UpdateUser.go
│   ├── middleware.go
│   └── jsonRes.go
│
├── models/
│   ├── db.go
│   └── models.go
│
├── index.html
├── home.html
├── navbar.html
├── profile.html
└── progress.html
```

---

## ⚙️ How It Works

1. User registers using **Signup API**
2. User logs in using **Login API**
3. After authentication, user can:

   * Add new tasks
   * View all tasks
   * Update task status/details
   * Delete tasks
4. User profile and task progress are displayed using HTML pages

---

## ▶️ How to Run the Project

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to project folder**

   ```bash
   cd todo
   ```

3. **Install dependencies**

   ```bash
   go mod tidy
   ```

4. **Run the server**

   ```bash
   go run main.go
   ```

5. Open browser and visit:

   ```
   http://localhost:8080
   ```

---

## 🔐 API Functionalities

* `POST /signup` – Register new user
* `POST /login` – Authenticate user
* `GET /tasks` – Fetch all tasks
* `POST /task` – Create task
* `PUT /task` – Update task
* `DELETE /task` – Delete task

---

## 🎯 Use Cases

* Daily task tracking
* Learning Go backend development
* Understanding REST APIs in Golang
* Beginner-friendly full-stack project

---

## 📌 Future Enhancements

* JWT-based authentication
* Password hashing & security improvements
* Responsive UI using CSS/JS
* Deployment on cloud (AWS/GCP)

---

## 👤 Author

**Dhruva G**
Aspiring Backend & Full-Stack Developer

---

If you want, I can also:

* 🔹 Make this **shorter for resume**
* 🔹 Add **screenshots section**
* 🔹 Convert it into **professional GitHub README**
* 🔹 Align it exactly with **your resume project description**

Just tell me 👍
