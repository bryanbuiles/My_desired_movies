# My desired movies - Api

My desired movies api is a movie api where you can search for movies and add them to a wishlist. Additionally it allows you to add authenticated users by means of JSON Web Token (JWT). Users will be able to store their desired movies in a list so they can watch them later.

The api also provides the basic services of a raw api such as get, post, update, delete for movies, users and wish list.

The api is still under development and the database is not populate with movies yet.

## Functionalities of this api:

- Create a user and movies
- Login with JWT (JSON Web Token).
- Create a list with my desired movies to watch later.
- Supports the functionality of a crud rest api (Get, Post, Put, Delete).

## Table of Content

- [Environment](#environment-and-requirements)
- [Run api locally](#Run-web-application-locally)
- [Endpoints](#Visit-our-web-site)
- [Folder Descriptions](#folder-descriptions)
- [Future improvements](#Future-improvements)
- [Bugs](#bugs)
- [Authors](#authors)
- [License](#license)

## Environment and requirements

This web-application was interpreted/tested on Ubuntu 20.04 LTS using go (version 1.15.6)

### General Requirements

- jwt: github.com/gofiber/jwt/v2
- fiber: github.com/gofiber/fiber/v2
- Database managment: github.com/lib/pq
- postgrepSQL
- Goland

## Run api locally

- Clone this repository: `git clone "https://github.com/bryanbuiles/My_desired_movies.git"`
- Access to My desired movies directory: `cd My_desired_movies/cmd`
- Update dependecies
  ```
  ~/My_desired_movies$ go mod tidy
  ```
- Run the script create_db in linux system
  ```
  ~/My_desired_movies$ ./create_db
  ```
- (Just for windows system) You have to create the database and the user manually and then run the posg.sql script for the database you created
- Run the api:
  ```
  ~/My_desired_movies$ go run main.go
  ```
- Request the endpoints with curl or postmant:

## Endpoints

### Movie endpoints:

- GET /movies - ALL movies
- GET /movies?director=Martin - Search for any parameter in movies
  Example:
  ```
  curl -X GET http://0.0.0.0:3001/movies?director=Martin
  ```
  Output:
  ```
  [
    {
        "id": "1b070506-342a-11eb-adc1-0242ac120002",
        "title": "Cape Fear",
        "caste": "Robert de Niro, Gregory Peck",
        "release_date": "1991-11-13T00:00:00Z",
        "genre": "suspense, drama",
        "director": "Martin scorses"
    }
  ]
  ```
- POST /movies - Create a movie
- PATCH /movies/movieID - Update a movie
- DELETE /movies/movieID - delete a movie

### User endpoints:

- GET /users - All users
- GET /users/userID - Get an user by id
- POST /users - Create an user
- POST /users/login - Login an user - provide a Bearer token
  Example:
  ```
  curl -X POST http://0.0.0.0:3001/users/login -H "Content-Type: application/json" -d '{"username": "bryan", "password": "123456"}'
  ```
  Output:
  ```
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9. eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjA3OTAyMzUxLCJzdWIiOiJmMjllYzU0Ni0wODQ3LTQ5ZGMtYTFiMy0xNGJmZDNkYjRkZTAifQ.9RkkSv9JxevTvxm-22vJig47woFqz5-0R3rxgbcpXZQ"
  }
  ```
- PATCH /users/userID - Update an user, needs the bearer token
- DELETE PATCH /users/userID - Delete an user, needs the bearer token

### wish list movies endpoints:

- GET /wishlist - all wish movies by user, needs the bearer token
- POST /wishlist - Add a new wish movie to the list, needs the bearer token

  Example:

  ```
  curl -X POST http://0.0.0.0:3001/wishlist -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjA3OTAxMzEzLCJzdWIiOiJmMjllYzU0Ni0wODQ3LTQ5ZGMtYTFiMy0xNGJmZDNkYjRkZTAifQ.gPdIREp7Ab2oljOkCfvApnNxal0XBNxNNfrybSbJa-I" -d '{"movie_id": "1b070506-342a-11eb-adc1-0242ac120002", "comment": "Me la vere despues"}'

  ```

  Output: `{"result":"Movie added to wish list"}`

- DELETE /wishlist/movieID - Delete a wish movie in the user wish list, needs the bearer token

## Folder descriptions

| Folder                       | Description                      |
| ---------------------------- | -------------------------------- |
| api                          | Contains all api files           |
| api/movies                   | Contains all movie files         |
| api/movies/models            | Contain all structs for movies   |
| api/movies/movie_gateway     | Contains all services for movies |
| api/movies/web_movie_handler | All handlers for movies          |
| api/users                    | Contains all users files         |
| api/users/models             | Contain all structs for users    |
| api/users/user_gateway       | All Services for users           |
| api/users/user_handler       | All handlres for useres          |
| cmd                          | Contain file main.go             |
| internal                     | Errors and database managment    |
| internal/database            | Database managment               |
| internal/logs                | Logs an errors                   |
| routes                       | Routes and endpoints for the api |
| scripts                      | Database scripts                 |

## Future improvements

- Add testing to the api - doing
- Docker
- Populate database with movies
- Add Frontend
- Deploy

## Bugs

No known bugs at this time.

## Authors

- Brayam Builes - [Github](https://github.com/bryanbuiles) / [Twitter](https://twitter.com/bryan_builes) / [Linkedin](https://www.linkedin.com/in/brayam-steven-builes-echavarria/)

## License

Apache-2.0 License.
