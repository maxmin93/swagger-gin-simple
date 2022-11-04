package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"swagger-gin-simple/models"
	"swagger-gin-simple/pkg/milestones"
)

// GetMilestones godoc
// @Summary Show all Milestones
// @Description get the Milestones with Tasks.
// @Tags milestones
// @Accept */*
// @Produce json
// @Success      200  {array}   models.Milestone
// @Router /api/v1/milestones [get]
func (api *RESTApiV1) GetMilestones(c *gin.Context) {
	milestones, err := milestones.GetService().GetAllMilestones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch milestones"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": milestones,
	})
}

// @Summary Update Milestone
// @Tags milestones
// @Accept */*
// @Produce json
// @Param        int   query      int  true  "int id" id(int)
// @Success      200  {object}  models.Milestone
// @Router /api/v1/milestones/{id} [put]
func (api *RESTApiV1) EditMilestone(c *gin.Context) {
	id := c.Param("id")
	var milestone models.Milestone
	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := milestones.GetService().EditMilestone(uint(idInt), milestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": milestone.ID,
	})
}

// @Summary Create a Milestone
// @Tags milestones
// @Accept */*
// @Produce json
// @Param request body models.Milestone true "new Milestone"
// @Success      200  {object}  models.Milestone
// @Router /api/v1/milestones [post]
func (api *RESTApiV1) AddMilestone(c *gin.Context) {

	var milestone models.Milestone
	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := milestones.GetService().CreateMilestone(&milestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": milestone.ID,
	})
}

// @Summary Delete the Milestone
// @Tags milestones
// @Accept */*
// @Produce json
// @Param int query int true "int id" id(int)
// @Success      200  {object}  models.Milestone
// @Router /api/v1/milestones [delete]
func (api *RESTApiV1) DeleteMilestone(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := milestones.GetService().DeleteMilestone(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
