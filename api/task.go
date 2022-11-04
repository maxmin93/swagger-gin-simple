package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"swagger-gin-simple/models"
	"swagger-gin-simple/pkg/tasks"
)

// GetTasks godoc
// @Summary Show all Tasks
// @Description get the Tasks with name, desc, done.
// @Tags tasks
// @Accept */*
// @Produce json
// @Success      200  {array}   models.Task
// @Router /api/v1/tasks [get]
func (api *RESTApiV1) GetTasks(c *gin.Context) {
	tasks, err := tasks.GetService().GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

// EditTasks godoc
// @Summary Update Task
// @Tags tasks
// @Accept */*
// @Produce json
// @Param        id   query      int  true  "int id" id(int)
// @Success      200  {object}  models.Task
// @Router /api/v1/tasks/{id} [put]
func (api *RESTApiV1) EditTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := tasks.GetService().EditTask(uint(idInt), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": task.ID,
	})
}

// AddTask godoc
// @Summary Create a Task
// @Tags tasks
// @Accept */*
// @Produce json
// @Param request body models.Task true "new Task"
// @Success      200  {object}  models.Task
// @Router /api/v1/tasks [post]
func (api *RESTApiV1) AddTask(c *gin.Context) {

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tasks.GetService().CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": task.ID,
	})
}

// DeleteTask godoc
// @Summary Delete the Task
// @Tags tasks
// @Accept */*
// @Produce json
// @Param        id   path      int  true  "int id" id(int)
// @Success      200  {id}  	int
// @Router /api/v1/tasks/{id} [delete]
func (api *RESTApiV1) DeleteTask(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := tasks.GetService().DeleteTask(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}