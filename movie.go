package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tidwall/gjson"
)

var db *gorm.DB
var e error

type Movie struct {
	ID        uint    `json:"-"`
	Name      string  `json:"name"`
	Plot      string  `json:"plot"`
	Year      int     `json:"year"`
	IMBDId    string  `json:"imbd_id"`
	IMDDScore float32 `json:"imbd_score"`
}

var moviesWait sync.WaitGroup

func main() {
	db, e = gorm.Open("sqlite3", "./example.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()

	//Start async getting movie plots
	executeGetMoviePlots()

	moviesWait.Wait()
}

// Gets all the movie plots in the current db and calls getMoviePlot async
func executeGetMoviePlots() {
	var movies []Movie
	if e := db.Find(&movies).Error; e != nil {
		fmt.Println(e)
	} else {
		for i := 0; i < len(movies); i++ {
			moviesWait.Add(1)
			go getMoviePlot(movies[i])
		}
	}
}

// Calls the api of omdbapi and gets the plot data then call update function
func getMoviePlot(movie Movie) {
	defer moviesWait.Done()
	resp, err := http.Get("http://www.omdbapi.com/?i=tt" + movie.IMBDId + "&apikey=fbbe832e")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyJSON := string(body)
	plotString := gjson.Get(bodyJSON, "Plot")
	movie.Plot = plotString.String()
	updateMovieByMovie(movie)
	fmt.Println(plotString)
}

// Update the movie by using the gorm orm
func updateMovieByMovie(movie Movie) {
	var mov Movie
	if e := db.Where("id = ?", movie.ID).First(&mov).Error; e != nil {
		fmt.Println(e)
	} else {
		mov.Plot = movie.Plot
		db.Save(&mov)
	}
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
	if e := db.Where("imbd_id = ?", id).First(&movie).Error; e != nil {
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
