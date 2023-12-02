package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // Importing Gorilla Mux for routing
)

// Movie struct to represent movie attributes
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct to represent director attributes
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // Slice to hold movie data

// Handler to get all movies and send JSON response
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Handler to delete a movie based on ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...) // Deleting the movie from the slice
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Handler to get a specific movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Handler to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)  // Decode request body into movie struct
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Generate a random ID
	movies = append(movies, movie)              // Add the new movie to the movies slice
	json.NewEncoder(w).Encode(movie)            // Respond with the created movie
}

// Handler to update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...) // Delete the existing movie
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the updated movie data from request body
	movie.ID = params["id"]                    // Set the ID from URL params
	movies = append(movies, movie)             // Append the updated movie to the movies slice
	json.NewEncoder(w).Encode(movie)           // Respond with the updated movie
}

func main() {
	r := mux.NewRouter() // Create a new router using Gorilla Mux

	movies = append(movies, Movie{ID: "1", Isbn: "123123", Title: "Harry Potter", Director: &Director{Firstname: "Mark", Lastname: "Badmash"}})
	movies = append(movies, Movie{ID: "2", Isbn: "876543", Title: "Chor machye shor", Director: &Director{Firstname: "Iron", Lastname: "Man"}})
	movies = append(movies, Movie{ID: "3", Isbn: "543221", Title: "Expandables", Director: &Director{Firstname: "Steph", Lastname: "Curry"}})
	movies = append(movies, Movie{ID: "4", Isbn: "456788", Title: "Captain America", Director: &Director{Firstname: "Luka", Lastname: "Doncic"}})
	movies = append(movies, Movie{ID: "5", Isbn: "135792", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Andy", Lastname: "Dufresne"}})
	movies = append(movies, Movie{ID: "6", Isbn: "987654", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "7", Isbn: "246810", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "8", Isbn: "112233", Title: "The Matrix", Director: &Director{Firstname: "Lana", Lastname: "Wachowski"}})
	movies = append(movies, Movie{ID: "9", Isbn: "998877", Title: "Forrest Gump", Director: &Director{Firstname: "Robert", Lastname: "Zemeckis"}})
	movies = append(movies, Movie{ID: "10", Isbn: "554433", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "11", Isbn: "223344", Title: "Fight Club", Director: &Director{Firstname: "David", Lastname: "Fincher"}})
	movies = append(movies, Movie{ID: "12", Isbn: "777888", Title: "Interstellar", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "13", Isbn: "888999", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "14", Isbn: "999000", Title: "Schindler's List", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: "15", Isbn: "000111", Title: "Goodfellas", Director: &Director{Firstname: "Martin", Lastname: "Scorsese"}})
	// Populate some initial movies in the movies slice
	// (just some dummy data for demonstration purposes)
	movies = append(movies, Movie{ID: "1", Isbn: "123123", Title: "Harry Potter", Director: &Director{Firstname: "Mark", Lastname: "Badmash"}})
	// ... (additional initial movie entries)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer)) // Serve static files

	// Route for the root path to render index.html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	// Define various API endpoints with corresponding handlers
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting at port 8082.......")
	err := http.ListenAndServe(":8082", r) // Start the server
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
