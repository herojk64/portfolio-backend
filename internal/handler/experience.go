package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio/internal/pkg/convert"
	"github.com/herojk64/portfolio/internal/pkg/response"
	"github.com/herojk64/portfolio/internal/service"
)

type ExperienceHandler struct {
	svc *service.ExperienceService
}

func NewExperienceHandler(svc *service.ExperienceService) *ExperienceHandler {
	return &ExperienceHandler{svc: svc}
}

type ExperienceRequest struct {
	Company     string  `json:"company"     binding:"required"`
	Role        string  `json:"role"        binding:"required"`
	Description *string `json:"description"`
	StartDate   string  `json:"start_date"  binding:"required"`
	EndDate     *string `json:"end_date"`
	Location    *string `json:"location"`
}

type ExperienceSkillRequest struct {
	SkillID string `json:"skill_id" binding:"required"`
}

// List godoc
// @Summary     List experiences
// @Tags        experience
// @Produce     json
// @Param       search  query  string false "Search company or role"
// @Param       company query  string false "Filter by company"
// @Param       role    query  string false "Filter by role"
// @Param       limit   query  int    false "Max results"       default(20)
// @Param       offset  query  int    false "Pagination offset" default(0)
// @Success     200 {object} response.Envelop
// @Router      /experience [get]
func (h *ExperienceHandler) List(c *gin.Context) {
	limit := int32(20)
	offset := int32(0)
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = int32(v)
		}
	}
	if o := c.Query("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil && v >= 0 {
			offset = int32(v)
		}
	}

	items, total, err := h.svc.List(
		context.Background(),
		c.Query("company"),
		c.Query("role"),
		c.Query("search"),
		limit, offset,
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"data": items, "total": total})
}

// Get godoc
// @Summary     Get experience by ID
// @Tags        experience
// @Produce     json
// @Param       id  path     string true "Experience UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Failure     404 {object} response.Envelop
// @Router      /experience/{id} [get]
func (h *ExperienceHandler) Get(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	exp, skills, err := h.svc.Get(context.Background(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "experience not found")
		return
	}
	response.Success(c, http.StatusOK, gin.H{"experience": exp, "skills": skills})
}

// Create godoc
// @Summary     Create an experience entry
// @Tags        experience
// @Accept      json
// @Produce     json
// @Param       body body     ExperienceRequest true "Experience data"
// @Success     201  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /experience [post]
func (h *ExperienceHandler) Create(c *gin.Context) {
	var req ExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	exp, err := h.svc.Create(context.Background(), service.ExperienceParams{
		Company:     req.Company,
		Role:        req.Role,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Location:    req.Location,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, exp)
}

// Update godoc
// @Summary     Update an experience entry
// @Tags        experience
// @Accept      json
// @Produce     json
// @Param       id   path     string            true "Experience UUID"
// @Param       body body     ExperienceRequest true "Experience data"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /experience/{id} [put]
func (h *ExperienceHandler) Update(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req ExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	exp, err := h.svc.Update(context.Background(), id, service.ExperienceParams{
		Company:     req.Company,
		Role:        req.Role,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Location:    req.Location,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, exp)
}

// Delete godoc
// @Summary     Delete an experience entry
// @Tags        experience
// @Produce     json
// @Param       id  path     string true "Experience UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Router      /experience/{id} [delete]
func (h *ExperienceHandler) Delete(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.Delete(context.Background(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}

// AddSkill godoc
// @Summary     Add a skill to an experience entry
// @Tags        experience
// @Accept      json
// @Produce     json
// @Param       id   path     string                 true "Experience UUID"
// @Param       body body     ExperienceSkillRequest true "Skill ID"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /experience/{id}/skills [post]
func (h *ExperienceHandler) AddSkill(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req ExperienceSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	skillID, err := convert.ParseUUID(req.SkillID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid skill_id")
		return
	}

	if err := h.svc.AddSkill(context.Background(), id, skillID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}

// RemoveSkill godoc
// @Summary     Remove a skill from an experience entry
// @Tags        experience
// @Produce     json
// @Param       id      path     string true "Experience UUID"
// @Param       skillId path     string true "Skill UUID"
// @Success     200     {object} response.Envelop
// @Failure     400     {object} response.Envelop
// @Router      /experience/{id}/skills/{skillId} [delete]
func (h *ExperienceHandler) RemoveSkill(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	skillID, err := convert.ParseUUID(c.Param("skillId"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid skill_id")
		return
	}

	if err := h.svc.RemoveSkill(context.Background(), id, skillID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}
