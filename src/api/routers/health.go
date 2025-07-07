package routers

import (
	"github.com/amirhasanpour/car-sale-management-wep-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/", handler.PostHealth)
	r.POST("/:id", handler.PostHealthById)
}