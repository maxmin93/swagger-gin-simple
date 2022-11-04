package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "swagger-gin-simple/docs/ginsimple"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func path(endpoint string) string {
	return fmt.Sprintf("/api/v1/%s", endpoint)
}

type RESTApiV1 struct {
	router *gin.Engine
}

func (api *RESTApiV1) Serve(addr string) error {
	return api.router.Run(addr)
}

func NewRESTApiV1() *RESTApiV1 {
	router := gin.Default()
	api := &RESTApiV1{
		router,
	}

    // health check
    router.GET("/", HealthCheck)

    // Swagger
    url := ginSwagger.URL("/swagger/doc.json")
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// projects
	router.POST(path("projects/:id"), api.EditProject)
	router.DELETE(path("projects/:id"), api.DeleteProject)
	router.GET(path("projects"), api.GetProjects)
	router.PUT(path("projects"), api.AddProject)

	// milestones
	router.POST(path("milestones/:id"), api.EditMilestone)
	router.DELETE(path("milestones/:id"), api.DeleteMilestone)
	router.GET(path("milestones"), api.GetMilestones)
	router.PUT(path("milestones"), api.AddMilestone)

	// tasks
	router.POST(path("tasks/:id"), api.EditTask)
	router.DELETE(path("tasks/:id"), api.DeleteTask)
	router.GET(path("tasks"), api.GetTasks)
	router.PUT(path("tasks"), api.AddTask)

	return api
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
    res := map[string]interface{}{
        "data": "Server is up and running",
    }

    c.JSON(http.StatusOK, res)
}
