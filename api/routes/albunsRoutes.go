package routes

import (
	controller "api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterAlbuns() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	router.GET("/", controller.GetInitial)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.PostAlbums)
	router.PUT("/albums/:id", controller.UpdateAlbum)
	router.DELETE("/albums/:id", controller.DeleteAlbum)
	router.Run(":8080")
}
