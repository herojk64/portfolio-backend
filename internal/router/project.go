package router

import (
	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio/internal/handler"
)

func ProjectRouter(rg *gin.RouterGroup, h *handler.ProjectHandler) {
	g := rg.Group("/projects")
	g.GET("", h.List)
	g.GET("/:id", h.Get)
	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.POST("/:id/skills", h.AddSkill)
	g.DELETE("/:id/skills/:skillId", h.RemoveSkill)
}
