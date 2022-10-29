package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creating a struct to sepcify the type of data we are looking for. In this small API build we will be passing JSON information to the API.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Album information that will be stored into the struct using variables.
var albums = []album{
	{ID: "1", Title: "Blow on this Blow fish", Artist: "Hooty and the Blowfish", Price: 59.99},
	{ID: "2", Title: "From Down Under", Artist: "Roo and the Aussie Crew", Price: 69.99},
	{ID: "3", Title: "Here we go again", Artist: "J.C. Clemens", Price: 39.99},
}

// Responds by getting a list of all albums and handles them in a JSON format
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	//Calling the bind JSON to append new information to the newAlbum variable
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	//This will add new albums to the album slice
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	//Locates an album using the ID in the get request.

	id := context.Param("id")

	//Looping over the albums to match the id and present that information to the client from the server.
	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost: 8080")
}
