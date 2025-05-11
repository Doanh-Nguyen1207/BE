package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// w stand for 'what do you write to?'
func (app *application) Home(w http.ResponseWriter, r *http.Request) { // ResponseWriter is where you write the final content you want to send to the client
	var payload = struct {
		Status  string `json:"status"`
		Message string `json: "message"`
		Version string `json:"version"`
	}{
		Status:  "OK",
		Message: "Hello world!",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movies

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movies{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		Runtime:     116,
		MPAARating:  "R",
		Description: "In the year 1536, a Scottish Highlander named Connor MacLeod is mortally wounded in battle. He is saved by a mysterious stranger, Ramirez, who reveals that he is immortal",
		Image:       "https://musicart.xboxlive.com/7/27b26300-0000-0000-0000-000000000002/504/image.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	rd, _ = time.Parse("2006-01-02", "1986-03-07")

	harrypotter := models.Movies{
		ID:          2,
		Title:       "Harry Potter",
		ReleaseDate: rd,
		Runtime:     120,
		MPAARating:  "PG-13",
		Description: "Harry Potter is a series of fantasy novels written by British author J.K. Rowling. The series follows the life and adventures of a young wizard, Harry Potter, and his friends Hermione Granger and Ron Weasley.",
		Image:       "https://img.posterstore.com/zoom/wb0101-8harrypotter-thephilosophersstoneno150x70.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, harrypotter)

	rd, _ = time.Parse("2006-01-02", "1986-03-07")

	harrypotter1 := models.Movies{
		ID:          3,
		Title:       "Harry Potte1",
		ReleaseDate: rd,
		Runtime:     120,
		MPAARating:  "PG-13",
		Description: "Harry Potter is a series of fantasy novels written by British author J.K. Rowling. The series follows the life and adventures of a young wizard, Harry Potter, and his friends Hermione Granger and Ron Weasley.",
		Image:       "https://static.wikia.nocookie.net/harrypotter/images/3/36/Harry-potter-films.png/revision/latest/scale-to-width-down/1200?cb=20110722151247",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, harrypotter1)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
