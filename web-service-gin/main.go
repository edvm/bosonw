package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Weather struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

const apiUrl = "https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s"

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/weather", getWeather)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getWeather(c *gin.Context) {
	lat, long := c.Query("lat"), c.Query("long")
	url := fmt.Sprintf(apiUrl, lat, long)

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	defer res.Body.Close()

	var weather Weather
	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	c.JSON(http.StatusOK, weather)
}
