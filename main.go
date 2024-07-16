package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Define a struct to represent a user
type User struct {
	Username string
	Password string
}

// Create a map to store users in-memory (not secure for production)
var users map[string]string = map[string]string{
	"admin": "password",
	"user":  "secret",
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Display login form
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Get username and password from form submission
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate credentials
	if validUser(username, password) {
		// Redirect to dashboard on successful login
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Display error message on failed login
	fmt.Fprintf(w, "Invalid username or password")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated (implement session management for real use)
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Display dashboard content
	fmt.Fprintf(w, "Welcome to the dashboard!")
}

func validUser(username string, password string) bool {
	// Check if username exists and password matches in the user map
	userPassword, ok := users[username]
	return ok && userPassword == password
}

func isAuthenticated(r *http.Request) bool {
	// Implement session management to check if user is logged in (omitted for simplicity)
	// This example always returns false, simulating no logged-in user
	return false
}
