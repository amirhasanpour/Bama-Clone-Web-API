package routers

import (
	"github.com/amirhasanpour/bama-clone-web-api/src/api/handlers"
	"github.com/amirhasanpour/bama-clone-web-api/src/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyCategoryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
