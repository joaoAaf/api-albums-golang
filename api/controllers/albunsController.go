package controllers

import (
	album "api/model"
	mongodb "api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAlbums(c *gin.Context) {
	if len(mongodb.FindAll()) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no albums found"})
		return
	}
	c.IndentedJSON(http.StatusOK, mongodb.FindAll())
}

func PostAlbums(c *gin.Context) {
	var newAlbum album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	c.IndentedJSON(http.StatusCreated, mongodb.InsertData(newAlbum).InsertedID)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	objId, errObjId := primitive.ObjectIDFromHex(id)
	if errObjId != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if len(mongodb.FindOne(objId)) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, mongodb.FindOne(objId))
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	objId, errObjId := primitive.ObjectIDFromHex(id)
	if errObjId != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var newAlbum album.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if mongodb.UpdateOne(objId, newAlbum) <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, mongodb.UpdateOne(objId, newAlbum))
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	objId, errObjId := primitive.ObjectIDFromHex(id)
	if errObjId != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if mongodb.DeleteOne(objId) <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, mongodb.DeleteOne(objId))
}
