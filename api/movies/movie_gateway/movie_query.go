package moviesgateway

// esta es la funcion para filtrar la busqueda de las peliculas por director, title or genre

import (
	"strings"

	"github.com/bryanbuiles/movie_suggester/api/movies/models"
)

func getMoviesQuery(filter models.MovieFilter) string {
	var (
		director, genre, title string
		clause                 bool   = false
		firstPartQuery         string = "select id, title, caste, release_date, genre, director from movie"
		stringBuilder                 = strings.Builder{} // estructura para almacenar las strings contruidas
	)
	stringBuilder.WriteString(firstPartQuery)

	switch {
	case filter.Director != "":
		director = "director like '%" + filter.Director + "%'"
		clause = true
	case filter.Genre != "":
		genre = "genre like '%" + filter.Genre + "%'"
		clause = true
	case filter.Title != "":
		title = "title like '%" + filter.Title + "%'"
		clause = true
	default:
		return stringBuilder.String() // retorna el string acomulado
	}
	if clause {
		var flag int
		stringBuilder.WriteString(" where ")
		if director != "" {
			stringBuilder.WriteString(director)
			flag = 1
		}
		if genre != "" {
			if flag == 1 {
				stringBuilder.WriteString(" or ")
			}
			stringBuilder.WriteString(genre)
			flag = 2
		}
		if title != "" {
			if flag == 1 || flag == 2 {
				stringBuilder.WriteString(" or ")
			}
			stringBuilder.WriteString(title)
		}
	}
	return stringBuilder.String()
}

// CreateMovieQuery querry to create user in DB
func CreateMovieQuery() string {
	return "INSERT INTO movie (id, title, caste, release_date, genre, director) VALUES ($1, $2, $3, TO_DATE($4, 'YYYY-MM-DD'), $5, $6)"
}

// DeleteMovieQuery To delete a Movie
func DeleteMovieQuery() string {
	return "DELETE FROM movie WHERE id = $1"
}

// UpdateMovieQuery querry to update movie
func UpdateMovieQuery(cmd models.Movie) string {
	var (
		title, caste, releaseDate, genre, director string
		firstPartQuery                             string = "UPDATE movie SET "
		flag                                       bool   = false
		stringBuilder                                     = strings.Builder{} // estructura para almacenar las strings contruidas
	)
	stringBuilder.WriteString(firstPartQuery)
	coma := ", "
	if cmd.Title != "" {
		title = "title = '" + cmd.Title + "'"
		flag = true
		stringBuilder.WriteString(title)
	}
	if cmd.Caste != "" {
		if flag == true {
			stringBuilder.WriteString(coma)
		}
		caste = "caste = '" + cmd.Caste + "'"
		flag = true
		stringBuilder.WriteString(caste)
	}
	if cmd.ReleaseDate != "" {
		if flag == true {
			stringBuilder.WriteString(coma)
		}
		releaseDate = "release_date = TO_DATE('" + cmd.ReleaseDate + "', 'YYYY-MM-DD')"
		flag = true
		stringBuilder.WriteString(releaseDate)
	}
	if cmd.Genre != "" {
		if flag == true {
			stringBuilder.WriteString(coma)
		}
		genre = "genre = '" + cmd.Genre + "'"
		flag = true
		stringBuilder.WriteString(genre)
	}
	if cmd.Director != "" {
		if flag == true {
			stringBuilder.WriteString(coma)
		}
		director = "director = '" + cmd.Director + "'"
		stringBuilder.WriteString(director)
	}
	stringBuilder.WriteString(" WHERE id = '" + cmd.ID + "'")
	return stringBuilder.String()
}
