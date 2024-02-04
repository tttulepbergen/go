package abrplus

import (
	"errors"
)

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}

var movies = []Movie{
	{
		ID:          "1",
		Title:       "Inception",
		Genre:       "Sci-Fi",
		ReleaseYear: 2010,
	},
	{
		ID:          "2",
		Title:       "The Shawshank Redemption",
		Genre:       "Drama",
		ReleaseYear: 1994,
	},
	{
		ID:          "3",
		Title:       "The Dark Knight",
		Genre:       "Action",
		ReleaseYear: 2008,
	},
	{
		ID:          "4",
		Title:       "Pulp Fiction",
		Genre:       "Crime",
		ReleaseYear: 1994,
	},
	{
		ID:          "5",
		Title:       "The Matrix",
		Genre:       "Sci-Fi",
		ReleaseYear: 1999,
	},
	{
		ID:          "6",
		Title:       "Forrest Gump",
		Genre:       "Drama",
		ReleaseYear: 1994,
	},
	{
		ID:          "7",
		Title:       "The Godfather",
		Genre:       "Crime",
		ReleaseYear: 1972,
	},
	{
		ID:          "8",
		Title:       "The Lord of the Rings: The Fellowship of the Ring",
		Genre:       "Adventure",
		ReleaseYear: 2001,
	},
	{
		ID:          "9",
		Title:       "Fight Club",
		Genre:       "Drama",
		ReleaseYear: 1999,
	},
	{
		ID:          "10",
		Title:       "The Silence of the Lambs",
		Genre:       "Thriller",
		ReleaseYear: 1991,
	},
	{
		ID:          "11",
		Title:       "Jurassic Park",
		Genre:       "Adventure",
		ReleaseYear: 1993,
	},
	{
		ID:          "12",
		Title:       "Titanic",
		Genre:       "Romance",
		ReleaseYear: 1997,
	},
}

func GetMovies() []Movie {
	return movies
}

func GetMovieByID(id string) (*Movie, error) {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("Movie not found")
}
