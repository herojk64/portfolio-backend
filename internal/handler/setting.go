package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/herojk64/portfolio-backend/internal/pkg/response"
	"github.com/herojk64/portfolio-backend/internal/service"
)

type SettingHandler struct {
	svc *service.SettingService
}

func NewSettingHandler(svc *service.SettingService) *SettingHandler {
	return &SettingHandler{svc: svc}
}

type SettingRequest struct {
	Key   string          `json:"key"   binding:"required"`
	Value json.RawMessage `json:"value" binding:"required" swaggertype:"object"`
}

type SettingUpdateRequest struct {
	Value json.RawMessage `json:"value" binding:"required" swaggertype:"object"`
}

type SettingResponse struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

func toSettingResponse(s sqlc.Setting) SettingResponse {
	return SettingResponse{Key: s.Key, Value: json.RawMessage(s.Value)}
}

// List godoc
// @Summary     List all settings
// @Tags        settings
// @Produce     json
// @Success     200 {object} response.Envelop
// @Router      /settings [get]
func (h *SettingHandler) List(c *gin.Context) {
	settings, err := h.svc.List(context.Background())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	out := make([]SettingResponse, len(settings))
	for i, s := range settings {
		out[i] = toSettingResponse(s)
	}
	response.Success(c, http.StatusOK, out)
}

// Get godoc
// @Summary     Get a setting by key
// @Description Keys in use: profile, social
// @Tags        settings
// @Produce     json
// @Param       key path     string true "Setting key (e.g. profile, social)"
// @Success     200 {object} response.Envelop
// @Failure     404 {object} response.Envelop
// @Router      /settings/{key} [get]
func (h *SettingHandler) Get(c *gin.Context) {
	setting, err := h.svc.Get(context.Background(), c.Param("key"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "setting not found")
		return
	}
	response.Success(c, http.StatusOK, toSettingResponse(setting))
}

// Upsert godoc
// @Summary     Create or update a setting
// @Description Creates the key if it does not exist, otherwise updates the value. Value is arbitrary JSON.
// @Tags        settings
// @Accept      json
// @Produce     json
// @Param       body body     SettingRequest true "Setting key and JSON value"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /settings [post]
func (h *SettingHandler) Upsert(c *gin.Context) {
	var req SettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	setting, err := h.svc.Upsert(context.Background(), req.Key, req.Value)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, toSettingResponse(setting))
}

// Update godoc
// @Summary     Update a setting's value
// @Tags        settings
// @Accept      json
// @Produce     json
// @Param       key  path     string              true "Setting key"
// @Param       body body     SettingUpdateRequest true "New JSON value"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /settings/{key} [put]
func (h *SettingHandler) Update(c *gin.Context) {
	var req SettingUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	setting, err := h.svc.Update(context.Background(), c.Param("key"), req.Value)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, toSettingResponse(setting))
}

// Delete godoc
// @Summary     Delete a setting
// @Tags        settings
// @Produce     json
// @Param       key path     string true "Setting key"
// @Success     200 {object} response.Envelop
// @Router      /settings/{key} [delete]
func (h *SettingHandler) Delete(c *gin.Context) {
	if err := h.svc.Delete(context.Background(), c.Param("key")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}
