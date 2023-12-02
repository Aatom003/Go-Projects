---

# Simple Form Handling Web App

This web application facilitates form submissions and demonstrates basic handling of HTTP POST requests in Go (Golang).

## Project Overview

The project sets up a server in Go to handle form submissions via HTTP POST requests. It employs basic form validation and logs incoming requests.

### Features

- **Form Submission**: Handles form submissions sent via HTTP POST method.
- **Form Validation**: Performs basic validation for the submitted form data.
- **HTTP Logging**: Logs incoming HTTP requests and serves static files.

## Project Setup

### Prerequisites

- Go (Golang) installed

### Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/simple-form-app.git
   cd WebServer
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

## Detailed Explanation

### File Structure

- `main.go`: Contains the main code for the application.
- `static/`: Directory to serve static files.

### Code Explanation

#### Dependencies

- `fmt`, `log`, `net/http`, `regexp`: Standard Go packages for I/O, logging, HTTP handling, and regular expressions.

#### Handlers

- `formHandler`: Handles form submissions via HTTP POST method, validates form data, and displays received data.
- `isValidFormData`: Performs basic validation for the submitted form data.
- `logRequest`: Middleware function to log incoming HTTP requests.

#### Functionality

- **POST `/form`**: Accepts form submissions, validates data, and displays received form data.

### Usage

- Access the form via the web interface by navigating to `http://localhost:8081`.
- Submit the form and observe the response displaying the submitted data.

## Future Improvements

- Enhance form validation with more robust checks.
- Implement database storage for submitted form data.
- Expand functionality to handle additional form elements and file uploads.

## Contributions

Feel free to contribute by opening issues, submitting pull requests, or suggesting enhancements.

## License

This project is licensed under the [MIT License](LICENSE).

---
