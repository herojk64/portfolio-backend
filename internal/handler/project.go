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

type ProjectHandler struct {
	svc *service.ProjectService
}

func NewProjectHandler(svc *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

type ProjectRequest struct {
	Title        string  `json:"title"         binding:"required"`
	Description  *string `json:"description"`
	RepoUrl      *string `json:"repo_url"`
	LiveUrl      *string `json:"live_url"`
	ImageUrl     *string `json:"image_url"`
	IsFeatured   bool    `json:"is_featured"`
	DisplayOrder int32   `json:"display_order"`
}

type ProjectSkillRequest struct {
	SkillID string `json:"skill_id" binding:"required"`
}

// List godoc
// @Summary     List projects
// @Tags        projects
// @Produce     json
// @Param       featured query  bool   false "Filter by featured status"
// @Param       search   query  string false "Search by title"
// @Param       limit    query  int    false "Max results"  default(20)
// @Param       offset   query  int    false "Pagination offset" default(0)
// @Success     200 {object} response.Envelop
// @Router      /projects [get]
func (h *ProjectHandler) List(c *gin.Context) {
	var featured *bool
	if f := c.Query("featured"); f != "" {
		v, err := strconv.ParseBool(f)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid featured param")
			return
		}
		featured = &v
	}

	limit := int32(20)
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

	projects, total, err := h.svc.List(context.Background(), featured, c.Query("search"), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, gin.H{"data": projects, "total": total})
}

// Get godoc
// @Summary     Get a project by ID
// @Tags        projects
// @Produce     json
// @Param       id  path string true "Project UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Failure     404 {object} response.Envelop
// @Router      /projects/{id} [get]
func (h *ProjectHandler) Get(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	project, skills, err := h.svc.Get(context.Background(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "project not found")
		return
	}
	response.Success(c, http.StatusOK, gin.H{"project": project, "skills": skills})
}

// Create godoc
// @Summary     Create a project
// @Tags        projects
// @Accept      json
// @Produce     json
// @Param       body body     ProjectRequest true "Project data"
// @Success     201  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.svc.Create(context.Background(), service.ProjectParams{
		Title:        req.Title,
		Description:  req.Description,
		RepoUrl:      req.RepoUrl,
		LiveUrl:      req.LiveUrl,
		ImageUrl:     req.ImageUrl,
		IsFeatured:   req.IsFeatured,
		DisplayOrder: req.DisplayOrder,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, project)
}

// Update godoc
// @Summary     Update a project
// @Tags        projects
// @Accept      json
// @Produce     json
// @Param       id   path     string         true "Project UUID"
// @Param       body body     ProjectRequest true "Project data"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /projects/{id} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	id, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.svc.Update(context.Background(), id, service.ProjectParams{
		Title:        req.Title,
		Description:  req.Description,
		RepoUrl:      req.RepoUrl,
		LiveUrl:      req.LiveUrl,
		ImageUrl:     req.ImageUrl,
		IsFeatured:   req.IsFeatured,
		DisplayOrder: req.DisplayOrder,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, project)
}

// Delete godoc
// @Summary     Delete a project
// @Tags        projects
// @Produce     json
// @Param       id  path     string true "Project UUID"
// @Success     200 {object} response.Envelop
// @Failure     400 {object} response.Envelop
// @Router      /projects/{id} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
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
// @Summary     Add a skill to a project
// @Tags        projects
// @Accept      json
// @Produce     json
// @Param       id   path     string              true "Project UUID"
// @Param       body body     ProjectSkillRequest true "Skill to add"
// @Success     200  {object} response.Envelop
// @Failure     400  {object} response.Envelop
// @Router      /projects/{id}/skills [post]
func (h *ProjectHandler) AddSkill(c *gin.Context) {
	projectID, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid project id")
		return
	}

	var req ProjectSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	skillID, err := convert.ParseUUID(req.SkillID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid skill id")
		return
	}

	if err := h.svc.AddSkill(context.Background(), projectID, skillID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}

// RemoveSkill godoc
// @Summary     Remove a skill from a project
// @Tags        projects
// @Produce     json
// @Param       id      path     string true "Project UUID"
// @Param       skillId path     string true "Skill UUID"
// @Success     200     {object} response.Envelop
// @Failure     400     {object} response.Envelop
// @Router      /projects/{id}/skills/{skillId} [delete]
func (h *ProjectHandler) RemoveSkill(c *gin.Context) {
	projectID, err := convert.ParseUUID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid project id")
		return
	}

	skillID, err := convert.ParseUUID(c.Param("skillId"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid skill id")
		return
	}

	if err := h.svc.RemoveSkill(context.Background(), projectID, skillID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, nil)
}
