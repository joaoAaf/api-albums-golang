package routes

import (
	"api/configs/middleware"
	controller "api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RouterAlbums() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	router.Use(middleware.PrometheusMiddleware)
	router.GET("/", controller.GetInitial)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.PostAlbums)
	router.PUT("/albums/:id", controller.UpdateAlbum)
	router.DELETE("/albums/:id", controller.DeleteAlbum)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":8080")
}
