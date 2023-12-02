---

# Movies CRUD App

This Movies CRUD (Create, Read, Update, Delete) application provides endpoints to manage a collection of movies using RESTful APIs.

## Project Overview

This project aims to create a simple CRUD application for managing a list of movies. It includes a backend written in Go (Golang) that offers RESTful endpoints to perform CRUD operations on movie entities.

### Features

- **Create Movie**: Add a new movie to the collection.
- **Read Movie**: Retrieve details of a specific movie or fetch all movies.
- **Update Movie**: Modify information about a particular movie.
- **Delete Movie**: Remove a movie from the collection.

## Project Setup

### Prerequisites

- Go (Golang) installed
- `gorilla/mux` package (install using `go get -u github.com/gorilla/mux`)

### Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/movies-crud-app.git
   cd CRUD
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

## Detailed Explanation

### File Structure

- `main.go`: Contains the main code for the application.
- `static/index.html`: The HTML file for the web interface (basic template).

### Code Explanation

#### Dependencies

- `github.com/gorilla/mux`: This package is used for routing in the application.

#### Movie and Director Structs

- `Movie` struct: Represents the attributes of a movie including ID, ISBN, Title, and Director details.
- `Director` struct: Represents the director's details including Firstname and Lastname.

#### Handlers

- `getMovies`: Handles HTTP GET request to fetch all movies.
- `deleteMovie`: Handles HTTP DELETE request to delete a movie by ID.
- `getMovie`: Handles HTTP GET request to get a movie by ID.
- `createMovie`: Handles HTTP POST request to create a new movie.
- `updateMovie`: Handles HTTP PUT request to update a movie by ID.

#### Functionality

- **GET `/movies`**: Retrieves all movies.
- **GET `/movies/{id}`**: Retrieves a specific movie by its ID.
- **POST `/movies`**: Creates a new movie entry.
- **PUT `/movies/{id}`**: Updates a movie entry by its ID.
- **DELETE `/movies/{id}`**: Deletes a movie entry by its ID.

### Usage

- Access the application through the web interface by navigating to `http://localhost:8082`.
- Utilize the provided API endpoints to perform CRUD operations on movies.

## Future Improvements

- Implement user authentication and authorization for secure endpoints.
- Enhance error handling and validation for user input.
- Incorporate a database for persistent storage.

## Contributions

Feel free to contribute by opening issues, submitting pull requests, or suggesting enhancements.

## License

This project is licensed under the [MIT License](LICENSE).

---
