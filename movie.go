package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"sync"

	"github.com/gin-contrib/cors"
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
	IMDDScore float64 `json:"imbd_score"`
}

var moviesWait sync.WaitGroup

func main() {
	db, e = gorm.Open("sqlite3", "./example.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()

	// db.AutoMigrate(&Movie{})

	r := gin.Default()
	r.Use(cors.Default())

	//FillMovies
	r.GET("/fill/movies", executeGetMoviePlots)
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

// Gets all the movie plots in the current db and calls getMoviePlot async
func executeGetMoviePlots(c *gin.Context) {
	var movies []Movie
	if e := db.Find(&movies).Error; e != nil {
		fmt.Println(e)
	} else {
		for i := 0; i < len(movies); i++ {
			moviesWait.Add(1)
			go getMoviePlot(movies[i], i)
			defer moviesWait.Done()
		}
	}
	moviesWait.Wait()
}

// Calls the api of omdbapi and gets the plot data then call update function
func getMoviePlot(movie Movie, i int) {
	fmt.Println(i)
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
	name := gjson.Get(bodyJSON, "Title")
	year := gjson.Get(bodyJSON, "Year")
	score := gjson.Get(bodyJSON, "imdbRating")
	movie.Plot = plotString.String()
	movie.Name = name.String()
	movie.Year = int(year.Int())
	movie.IMDDScore = score.Float()
	fmt.Println(movie)
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
		mov.Name = movie.Name
		mov.Year = movie.Year
		mov.IMDDScore = movie.IMDDScore
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

	getMoviePlot(movie, 0)
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
