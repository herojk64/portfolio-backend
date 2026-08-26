package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/herojk64/portfolio/docs"
	"github.com/herojk64/portfolio/internal/database/sqlc"
	"github.com/herojk64/portfolio/internal/handler"
	"github.com/herojk64/portfolio/internal/pkg/response"
	"github.com/herojk64/portfolio/internal/router"
	"github.com/herojk64/portfolio/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
)

func Handle(r *gin.Engine, q *sqlc.Queries) {
	version := r.Group("/api/v1")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		response.Success(c, http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/form", func(c *gin.Context) {
		token := csrf.GetToken(c)
		response.Success(c, http.StatusOK, gin.H{"csrf_token": token})
	})

	userSvc := service.NewUserService(q)
	userHandler := handler.NewUserHandler(userSvc)
	router.UserRouter(version, userHandler)

	projectSvc := service.NewProjectService(q)
	projectHandler := handler.NewProjectHandler(projectSvc)
	router.ProjectRouter(version, projectHandler)

	skillSvc := service.NewSkillService(q)
	skillHandler := handler.NewSkillHandler(skillSvc)
	router.SkillRouter(version, skillHandler)

	settingSvc := service.NewSettingService(q)
	settingHandler := handler.NewSettingHandler(settingSvc)
	router.SettingRouter(version, settingHandler)

	experienceSvc := service.NewExperienceService(q)
	experienceHandler := handler.NewExperienceHandler(experienceSvc)
	router.ExperienceRouter(version, experienceHandler)
}
