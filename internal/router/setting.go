package router

import (
	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio/internal/handler"
)

func SettingRouter(rg *gin.RouterGroup, h *handler.SettingHandler) {
	g := rg.Group("/settings")
	g.GET("", h.List)
	g.GET("/:key", h.Get)
	g.POST("", h.Upsert)
	g.PUT("/:key", h.Update)
	g.DELETE("/:key", h.Delete)
}
