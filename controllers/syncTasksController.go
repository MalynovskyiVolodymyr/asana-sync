package controllers

import (
	"asanaSync/common"
	"asanaSync/models"
	"asanaSync/services"
	"fmt"
	"strconv"
)

func SyncTasks(mittingId, projectId int64) {
	//var err error
	fmt.Println("*************")
	fmt.Println("FETCHING TASKS FROM TRACTION TOOLS")
	tractionTasks := services.GetTasksByMitting(mittingId)

	fmt.Println("*************")
	fmt.Println("LENGTH OF TASKS AFTER FETCH")
	fmt.Println(len(tractionTasks))
	fmt.Println("*************")
	tasks := GetAllTractionTasks()

	if len(tasks) == 0 {
		SaveTractionTasks(tractionTasks)
		fmt.Println("LENGTH AFTER SAVING")
		fmt.Println(len(tractionTasks))
		fmt.Println("----------------------------------------")
	} else {
		if len(tasks) >= len(tractionTasks) {
			fmt.Println("LENGTH AFTER FILTERING")
			tractionTasks = filteringTasks(tasks, tractionTasks)
			fmt.Println("***************************")
			fmt.Println(len(tractionTasks))

		} else {
			tractionTasks = filteringTasks(tractionTasks, tasks)
			fmt.Println("LENGTH AFTER FILTERING")
			fmt.Println(len(tractionTasks))
			fmt.Println("***************************")
			fmt.Println(len(tractionTasks))
		}
	}

	if len(tractionTasks) != 0 {
		asanaUsers := GetAllAsanaUsers()
		tractionUsers := GetAllTractionUsers()
		SaveTractionTasks(tractionTasks)
		saveToAsana(projectId, tractionTasks, asanaUsers, tractionUsers)
	}

	//
	fmt.Println("*************")
	SyncTasks(mittingId, projectId)
}

func saveToAsana(projectId int64, tasks []models.TractionTaskModel, asanaUsers []models.AsanaUser, tractionUsers []models.TractionUserModel) {

	fmt.Println("SAVING DATA TO ASANA ITEMS LEFT => " + strconv.Itoa(len(tasks)))
	if len(tasks) == 0 {
		return
	}
	services.AsanaCreateTask(tasks[0].DetailsURL, tasks[0].Name, findAssignee(tasks[0].Owner.ID, asanaUsers, tractionUsers), tasks[0].DueDate, projectId)
	tasks = common.RemoveTractionTaskFromArray(tasks, 0)
	saveToAsana(projectId, tasks, asanaUsers, tractionUsers)
}

func findAssignee(ownerID int64, asanaUsers []models.AsanaUser, tractionUsers []models.TractionUserModel) string {

	for _, value := range tractionUsers {
		if ownerID == value.UserID {
			for _, value2 := range asanaUsers {
				if value.Email == value2.Email {
					return value2.Gid
				}
			}
		}
	}

	return common.USER_WAN_NOT_FOND_ERROR
}
