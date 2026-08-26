package router

import (
	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio-backend/internal/handler"
)

func UserRouter(rg *gin.RouterGroup, h *handler.UserHandler) {
	rg.GET("/users", h.List)
	rg.POST("/users", func(c *gin.Context) {})
}
