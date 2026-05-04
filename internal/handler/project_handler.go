package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"caselog/internal/model"
	"caselog/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *ProjectHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"Title": "案件登録 | CaseLog",
	})
}

func (h *ProjectHandler) Create(c *gin.Context) {
	unitPrice, err := strconv.Atoi(c.PostForm("unit_price"))
	if err != nil {
		unitPrice = 0
	}

	project := model.Project{
		Title:      c.PostForm("title"),
		Company:    c.PostForm("company"),
		MainSkill:  c.PostForm("main_skill"),
		UnitPrice:  unitPrice,
		RemoteType: c.PostForm("remote_type"),
		TeamSize:   c.PostForm("team_size"),
		Status:     c.PostForm("status"),
		Notes:      c.PostForm("notes"),
	}

	if project.Status == "" {
		project.Status = "considering"
	}

	if err := h.Repo.Create(project); err != nil {
		c.String(http.StatusInternalServerError, "案件の登録に失敗しました")
		return
	}

	c.Redirect(http.StatusFound, "/")
}

type ProjectHandler struct {
	Repo *repository.ProjectRepository
}

func NewProjectHandler(repo *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{Repo: repo}
}

func (h *ProjectHandler) Index(c *gin.Context) {
	projects, err := h.Repo.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "案件一覧の取得に失敗しました")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":       "CaseLog",
		"Projects":    projects,
		"StatusLabel": model.StatusLabel,
	})
}

func (h *ProjectHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "不正なIDです")
		return
	}

	project, err := h.Repo.FindByID(id)
	if err != nil {
		c.String(http.StatusNotFound, "案件が見つかりませんでした")
		return
	}

	c.HTML(http.StatusOK, "show.html", gin.H{
		"Title":       project.Title + " | CaseLog",
		"Project":     project,
		"StatusLabel": model.StatusLabel,
	})
}

func (h *ProjectHandler) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid project ID")
		return
	}

	project, err := h.Repo.FindByID(id)
	if err != nil {
		c.String(http.StatusNotFound, "Project not found")
		return
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"project": project,
	})
}

func (h *ProjectHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid project ID")
		return
	}

	unitPrice, err := strconv.Atoi(c.PostForm("unit_price"))
	if err != nil {
		unitPrice = 0
	}

	project := model.Project{
		ID:         id,
		Title:      c.PostForm("title"),
		Company:    c.PostForm("company"),
		MainSkill:  c.PostForm("main_skill"),
		UnitPrice:  unitPrice,
		RemoteType: c.PostForm("remote_type"),
		TeamSize:   c.PostForm("team_size"),
		Status:     c.PostForm("status"),
		Notes:      c.PostForm("notes"),
	}

	if err := h.Repo.Update(project); err != nil {
		c.String(http.StatusInternalServerError, "Failed to update project")
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d", id))
}