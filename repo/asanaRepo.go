package repo

import (
	"asanaSync/models"

	"gopkg.in/mgo.v2"
)

// AsanaRepository asana repo obj
type AsanaRepository struct {
	C *mgo.Collection
}

// GetAsanaUsers asana users
func (repo *AsanaRepository) GetAsanaUsers() []models.AsanaUser {
	var users []models.AsanaUser
	result := models.AsanaUser{}

	iter := repo.C.Find(nil).Iter()

	for iter.Next(&result) {
		users = append(users, result)
	}

	return users
}

// GetAsanaAllProjects asana projects
func (repo *AsanaRepository) GetAsanaAllProjects() []models.AsanaProject {
	var projects []models.AsanaProject
	result := models.AsanaProject{}

	iter := repo.C.Find(nil).Iter()

	for iter.Next(&result) {
		projects = append(projects, result)
	}

	return projects
}
