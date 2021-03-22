package models

// Movie
type Movie struct {
	ID               int     `json:"id"`
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// MovieDetail
type MovieDetail struct {
	ID               int      `json:"id"`
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	Genres           []Genres `json:"genres"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date"`
	Title            string   `json:"title"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// Genres
type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
