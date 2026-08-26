package router

import (
	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio-backend/internal/handler"
)

func SkillRouter(rg *gin.RouterGroup, h *handler.SkillHandler) {
	g := rg.Group("/skills")
	g.GET("", h.List)
	g.GET("/:id", h.Get)
	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
