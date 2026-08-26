package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/herojk64/portfolio-backend/internal/pkg/response"
	"github.com/herojk64/portfolio-backend/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.svc.ListUsers(context.Background(), sqlc.SelectUsersParams{})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, users)
}
