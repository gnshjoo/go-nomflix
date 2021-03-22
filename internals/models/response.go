package models

type MovieResponse struct {
	Dates        []string
	Page         int
	Results      []Movie
	TotalPages   int
	TotalResults int
}

type TVResponse struct {
	Dates        []string
	Page         int
	Results      []TV
	TotalPages   int
	TotalResults int
}

