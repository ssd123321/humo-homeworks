package main

type Movie struct {
	title  string
	rating float64
}
type MovieList struct {
	Movies []Movie
}

func (m *MovieList) AddMovie(title string, rating float64) {
	m.Movies = append(m.Movies, Movie{title, rating})
}
func (m MovieList) AverageRating() float64 {
	var sum float64
	for _, movie := range m.Movies {
		sum += movie.rating
	}
	return sum / float64(len(m.Movies))
}
