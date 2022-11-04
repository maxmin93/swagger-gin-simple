package seeds

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"

	"swagger-gin-simple/models"
	"swagger-gin-simple/pkg/milestones"
	"swagger-gin-simple/pkg/projects"
	"swagger-gin-simple/pkg/tasks"
)

type Seeds struct {
	Tasks      []models.Task
	Projects   []models.Project
	Milestones []models.Milestone
}

func RunSeeds(seedFilePath string) error {

	taskService := tasks.GetService()
	projectService := projects.GetService()
	milestoneService := milestones.GetService()

	dbTasks, err := taskService.GetAllTasks()
	if err != nil {
		return err
	}
	if len(dbTasks) > 0 {
		logrus.Info("Tasks already exist, skipping all migrations")
		return nil
	}

	seedFile, err := os.Open(seedFilePath)
	if err != nil {
		return err
	}
	defer seedFile.Close()
	var data Seeds
	if err := json.NewDecoder(seedFile).Decode(&data); err != nil {
		return err
	}

	for _, u := range data.Projects {
		if err := projectService.CreateProject(&u); err != nil {
			return err
		}
	}

	for _, r := range data.Milestones {
		if err := milestoneService.CreateMilestone(&r); err != nil {
			return err
		}
	}

	for _, p := range data.Tasks {
		if err := taskService.CreateTask(&p); err != nil {
			return err
		}
	}

	return nil
}