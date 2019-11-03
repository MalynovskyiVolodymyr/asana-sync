package controllers

import (
	"asanaSync/common"
	"asanaSync/data"
	"asanaSync/models"
	"asanaSync/repo"
	"fmt"
	"log"
	"strconv"
)

func GetAllTractionUsers() []models.TractionUserModel {
	context := data.NewContext()

	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollUserTraction)
	repo := repo.TractionRepository{collection}

	return repo.GetTractionUsers()
}

func GetAllTractionTasks() []models.TractionTaskModel {
	context := data.NewContext()

	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollTaskTraction)
	repo := repo.TractionRepository{collection}

	return repo.GetAllTasks()
}

func GetAllTractionMittings() []models.TractionMittingModel {
	context := data.NewContext()

	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollMittingTraction)
	repo := repo.TractionRepository{collection}

	return repo.GetAllMittings()
}

func SaveTractionTasks(tasks []models.TractionTaskModel) {
	if len(tasks) == 0 {
		return
	}
	fmt.Println("in save tasks ==> " + strconv.Itoa(len(tasks)))
	context := data.NewContext()
	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollTaskTraction)
	repo := repo.TractionRepository{collection}

	err := repo.SaveTask(tasks[0])

	if err != nil {
		log.Fatal(err)
	}
	tasks = common.RemoveTractionTaskFromArray(tasks, 0)
	SaveTractionTasks(tasks)
}

func filteringTasks(savedTasks, newTasks []models.TractionTaskModel) []models.TractionTaskModel {

	for _, value := range savedTasks {
		for k, value2 := range newTasks {
			if value.ID == value2.ID {
				newTasks = common.RemoveTractionTaskFromArray(newTasks, k)
			}
		}
	}

	return newTasks
}
