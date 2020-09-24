package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var e error

type Movie struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Year      int     `json:"year"`
	IMBDId    int     `json:"imbd_id"`
	IMDDScore float32 `json:"imbd_score"`
}

func main() {
	db, e = gorm.Open("sqlite3", "./example.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()

	db.AutoMigrate(&Movie{})

	r := gin.Default()
	// Get movies
	r.GET("/movies", getMovies)
	// Get movies by id
	r.GET("/movies/:id", getMovieByid)
	// Insert new movie
	r.POST("/movies", insertMovie)
	// Update movie
	r.PUT("/movies/:id", updateMovie)
	// Delete movie
	r.DELETE("/movies/:id", deleteMovie)
	r.Run(":1991")
}

// Get movies
func getMovies(c *gin.Context) {
	var movies []Movie
	if e := db.Find(&movies).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, movies)
	}
}

// Get movie by id
func getMovieByid(c *gin.Context) {
	var movie Movie
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&movie).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, movie)
	}
}

// Insert new movie
func insertMovie(c *gin.Context) {
	var movie Movie
	c.BindJSON(&movie)
	db.Create(&movie)
	c.JSON(200, movie)
}

// Update movie
func updateMovie(c *gin.Context) {
	var movie Movie
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&movie).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.BindJSON(&movie)
		db.Save(&movie)
		c.JSON(200, movie)
	}
}

// Delete movie
func deleteMovie(c *gin.Context) {
	var movie Movie
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&movie)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
