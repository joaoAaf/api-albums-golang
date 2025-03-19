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
	if len(mongodb.FindAll()) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no albums found"})
		return
	}
	c.IndentedJSON(http.StatusOK, mongodb.FindAll())
}

func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	c.IndentedJSON(http.StatusCreated, mongodb.InsertData(newAlbum).InsertedID)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	objId, errObjId := primitive.ObjectIDFromHex(id)
	if errObjId != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if mongodb.FindOne(objId) == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, mongodb.FindOne(objId))
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}
