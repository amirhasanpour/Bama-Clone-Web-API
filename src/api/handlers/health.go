package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {

}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "working!")
}

func (h *HealthHandler) PostHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "post working!")
}

func (h *HealthHandler) PostHealthById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("post working by id: %s!", id))
}