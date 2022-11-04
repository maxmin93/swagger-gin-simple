package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"swagger-gin-simple/models"
	"swagger-gin-simple/pkg/projects"
)

// GetProjects godoc
// @Summary Show all Projects
// @Description get the Projects with Milestones.
// @Tags projects
// @Accept */*
// @Produce json
// @Success      200  {array}   models.Project
// @Router /api/v1/projects [get]
func (api *RESTApiV1) GetProjects(c *gin.Context) {
	projects, err := projects.GetService().GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})
}

// @Summary Update Project
// @Tags projects
// @Accept */*
// @Produce json
// @Param int query int true "int id" id(int)
// @Success      200  {object}  models.Project
// @Router /api/v1/projects [put]
func (api *RESTApiV1) EditProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := projects.GetService().EditProject(uint(idInt), project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": project.ID,
	})
}

// @Summary Create a Project
// @Tags projects
// @Accept */*
// @Produce json
// @Param request body models.Project true "new Project"
// @Success      200  {object}  models.Project
// @Router /api/v1/projects [post]
func (api *RESTApiV1) AddProject(c *gin.Context) {

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := projects.GetService().CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": project.ID,
	})
}

// @Summary Delete the project
// @Tags projects
// @Accept */*
// @Produce json
// @Param int query int true "int id" id(int)
// @Success      200  {object}  models.Project
// @Router /api/v1/projects [delete]
func (api *RESTApiV1) DeleteProject(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := projects.GetService().DeleteProject(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
