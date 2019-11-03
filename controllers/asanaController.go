package controllers

import (
	"asanaSync/common"
	"asanaSync/data"
	"asanaSync/models"
	"asanaSync/repo"
)

func GetAllAsanaUsers() []models.AsanaUser {
	context := data.NewContext()

	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollUserAsana)
	repo := repo.AsanaRepository{collection}

	return repo.GetAsanaUsers()
}

func GetAllAsanaProjects() []models.AsanaProject {
	context := data.NewContext()

	defer context.Close()

	collection := context.DBCollection(common.AppConfig.CollProjectAsana)
	repo := repo.AsanaRepository{collection}

	return repo.GetAsanaAllProjects()
}
