package main // Indicates that this Go file belongs to the 'main' package

import (
	"fmt"      // Importing 'fmt' for formatted I/O
	"log"      // Importing 'log' for logging
	"net/http" // Importing 'net/http' for HTTP handling
	"regexp"   // Importing 'regexp' for regular expressions
	// Importing 'strings' for string manipulation
)

// formHandler handles form submissions via HTTP POST method
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not POST; return an error if not allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the form data from the request
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusInternalServerError)
		log.Println("Form Parse Error:", err) // Log form parsing errors
		return
	}

	// Set the content type of the response to HTML for proper rendering
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Retrieve form values for 'name', 'email', and 'message'
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Validation: Check if any of the form fields are empty or invalid
	if !isValidFormData(name, email, message) {
		http.Error(w, "<h2>Invalid or incomplete data submitted</h2>", http.StatusBadRequest)
		return
	}

	// Display the received form data in the response
	fmt.Fprintf(w, "<h2>Form Data Received:</h2>")
	fmt.Fprintf(w, "<p><strong>Name:</strong> %s</p>", name)
	fmt.Fprintf(w, "<p><strong>Email:</strong> %s</p>", email)
	fmt.Fprintf(w, "<p><strong>Message:</strong> %s</p>", message)
}

// Function to validate form data (basic validation for demonstration purposes)
func isValidFormData(name, email, message string) bool {
	// Email validation using a simple regular expression
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if len(name) == 0 || len(message) == 0 || !emailRegex.MatchString(email) {
		return false
	}
	return true
}

func main() {
	// Create a file server to serve static files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle requests to the root path ("/") by serving files from the file server
	http.Handle("/", fileServer)

	// Handle requests to "/form" by invoking the 'formHandler' function for POST requests
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is starting on port 8081...")

	// Custom error pages for different HTTP status codes
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "<h2>Custom Error Page</h2>", http.StatusBadRequest)
	})

	// Start the HTTP server on port 8081 and log any errors
	err := http.ListenAndServe(":8081", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Middleware function to log HTTP requests
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
