package app

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio/internal/config"
	"github.com/herojk64/portfolio/internal/pkg/response"
	"github.com/utrack/gin-csrf"
)

func New(cfg *config.Config) *gin.Engine {
	store := cookie.NewStore([]byte(cfg.Session.Secret))

	router := gin.New()

	router.Use(
		cors.New(cors.Config{
			AllowOrigins:     cfg.App.AllowedHost,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		RateLimiter(),
		sessions.Sessions("session", store),
		csrf.Middleware(csrf.Options{
			Secret: cfg.CSRF.Secret,
			ErrorFunc: func(c *gin.Context) {
				response.Error(c, http.StatusForbidden, "CSRF token mismatch")
			},
		}),
		gin.Logger(),
		gin.Recovery(),
	)

	return router
}
