package app

import (
	"asanaSync/common"
	"asanaSync/controllers"
)

// Start - app start
func Start() {
	common.SeedAsanaAndTracUsers()
	projects := controllers.GetAllAsanaProjects()
	mittings := controllers.GetAllTractionMittings()
	mittingAndPrj := common.PrepareMittingsAndProjects(mittings, projects)

	for _, value := range mittingAndPrj {
		go controllers.SyncTasks(value.MittingID, value.ProjectID)
	}
}
