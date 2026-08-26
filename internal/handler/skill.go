package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herojk64/portfolio-backend/internal/pkg/convert"
	"github.com/herojk64/portfolio-backend/internal/pkg/response"
	"github.com/herojk64/portfolio-backend/internal/service"
)

type SkillHandler struct {
	svc *service.SkillService
}

func NewSkillHandler(svc *service.SkillService) *SkillHandler {
	return &SkillHandler{svc: svc}
}

type SkillRequest struct {
	Name         string  `json:"name"          binding:"required"`
	Category     *string `json:"category"`
	Proficiency  *string `json:"proficiency"`
	Icon         *string `json:"icon"`
	DisplayOrder int32   `json:"display_order"`
}

// List godoc
// @Summary     List skills
// @Tags        skills
// @Produce     json
// @Param       category query  string false "Filter by category (Backend, Frontend, DevOps, Database, Scripting)"
// @Param       search   query  string false "Search by name"
// @Param       limit    query  int    false "Max results"      default(50)
// @Param       offset   query  int    false "Pagination offset" default(0)
// @Success     200 {object} response.Envelop
// @Router      /skills [get]
func (h *SkillHandler) List(c *gin.Context) {
	var category *string
	if cat := c.Query("category"); cat != "" {
		category = &cat
	}

	limit := int32(50)
	offset := int32(0)
	if l := c.Query("limit"); l != "" {
		v, _ := strconv.Atoi(l)
		if v > 0 {
			limit = int32(v)
		}
	}
	if o := c.Query("offset"); o != "" {
		v, _ := strconv.Atoi(o)
		if v >= 0 {
			offset = int32(v)
		}
	}

	skills, total, err := h.svc.List(context.Background(), category, c.Query("search"), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"data": skills, "total": total})
}

// Get godoc
// @Summary     Get a skill by ID
// @Tags        skills
// @Produce     json
// @Param       id  path     string true "Skill UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Failure     404 {object} response.Envelop
// @Router      /skills/{id} [get]
func (h *SkillHandler) Get(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	skill, err := h.svc.Get(context.Background(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "skill not found")
		return
	}
	response.Success(c, http.StatusOK, skill)
}

// Create godoc
// @Summary     Create a skill
// @Tags        skills
// @Accept      json
// @Produce     json
// @Param       body body     SkillRequest true "Skill data"
// @Success     201  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /skills [post]
func (h *SkillHandler) Create(c *gin.Context) {
	var req SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	skill, err := h.svc.Create(context.Background(), service.SkillParams{
		Name:         req.Name,
		Category:     req.Category,
		Proficiency:  req.Proficiency,
		Icon:         req.Icon,
		DisplayOrder: req.DisplayOrder,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, skill)
}

// Update godoc
// @Summary     Update a skill
// @Tags        skills
// @Accept      json
// @Produce     json
// @Param       id   path     string       true "Skill UUID"
// @Param       body body     SkillRequest true "Skill data"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /skills/{id} [put]
func (h *SkillHandler) Update(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	skill, err := h.svc.Update(context.Background(), id, service.SkillParams{
		Name:         req.Name,
		Category:     req.Category,
		Proficiency:  req.Proficiency,
		Icon:         req.Icon,
		DisplayOrder: req.DisplayOrder,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, skill)
}

// Delete godoc
// @Summary     Delete a skill
// @Tags        skills
// @Produce     json
// @Param       id  path     string true "Skill UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Router      /skills/{id} [delete]
func (h *SkillHandler) Delete(c *gin.Context) {
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
