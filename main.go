package main

import (
    "net/http"
	"fmt"
    "github.com/gin-gonic/gin"
)


type Album struct {
    ID     string     `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// album represents data about a record album.
var albums = []Album {
	{ID: "1", Title: "Take my breath away", Artist: "Berlin", Price: 12.99},
	{ID: "2", Title: "Paint the twon red", Artist: "Doja Cat", Price: 23.99},
	{ID: "3", Title: "Money trees", Artist: "Kenderick Lamar", Price:45.99},
	{ID: "4", Title: "Toy soldiers", Artist: "Eminem", Price: 69.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbum)

	router.Run("localhost:8080")	
}

// getting a list of all albums  as a json response using gin
func getAlbums(c *gin.Context) {
    fmt.Printf("Albums: %+v\n", albums)
    c.IndentedJSON(http.StatusOK, albums)
}

// making a POST reguest to create new Album
func postAlbums (c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}


// Getting an album by the ID if the album

func getAlbum(c *gin.Context){
	id := c.Param("id")

	for _, item := range albums {
		if (item.ID == id){	
			c.IndentedJSON(http.StatusOK, item)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found, sorry!"})
}