package api

// esta es la funcion para filtrar la busqueda de las peliculas por director, title or genre

import (
	"strings"
)

func getMoviesQuery(filter MovieFilter) string {
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

// CreateUserQuery ...
func CreateUserQuery() string {
	return "insert into users (id, username, password) values ($1, $2, $3)"
}
