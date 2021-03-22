package models

// TV
type TV struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	OriginalLanguage string  `json:"original_language"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// TVDetail
type TVDetail struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	BackdropPath     string  `json:"backdrop_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}
