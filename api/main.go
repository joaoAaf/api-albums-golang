package main

import (
	mongodb "api/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title"`
	Artist string             `bson:"artist"`
	Price  float64            `bson:"price"`
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, mongodb.FindAll())
}

func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	c.JSON(http.StatusCreated, mongodb.InsertData(newAlbum).InsertedID)
}

/*
	func getAlbumByID(c *gin.Context) {
		id := c.Param("id")
		for _, a := range albums {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
*/

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	/*router.GET("/albums/:id", getAlbumByID)*/
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}
