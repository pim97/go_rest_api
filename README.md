**How to run**

1. unzip the results.zip (most likely already done)

2. get the following repo's
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite

3. go run movie.go

4. server setup at port 1991

5. database sqlite example.db is already included

7. POST http://localhost:1991/movies

Add following body
{"name": "hackathon movie`", "year": 1964, "imbd_id": 10, "imbd_score": 5.5}

8. GET by ID
GET http://localhost:1991/movies/10

Example Response
Response
{
    "id": 4,
    "name": "hackathon movie`",
    "year": 1964,
    "imbd_id": 10,
    "imbd_score": 5.5
}

9. finished