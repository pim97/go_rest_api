**How to run**
go run movie.go

Database example.db is already included

**Endpoint URL**
http://localhost:1991/movies

**Example GET**
GET http://localhost:1991/movies

**Example POST**
POST http://localhost:1991/movies

Add following body
{"name": "hackathon movie`", "year": 1964, "imbd_id": 4, "imbd_score": 5.5}

**Example GET by ID**
GET http://localhost:1991/movies/4

Response
{
    "id": 4,
    "name": "hackathon movie`",
    "year": 1964,
    "imbd_id": 4,
    "imbd_score": 5.5
}

**All Endpoints**

r.GET("/movies", getMovies)
// Get movies by id
r.GET("/movies/:id", getMovieByid)
// Insert new movie
r.POST("/movies", insertMovie)
// Update movie
r.PUT("/movies/:id", updateMovie)
// Delete movie
r.DELETE("/movies/:id", deleteMovie)

