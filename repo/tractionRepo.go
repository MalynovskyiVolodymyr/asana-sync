package repo

import (
	"asanaSync/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TractionRepository tr
type TractionRepository struct {
	C *mgo.Collection
}

// GetTractionUsers tr users
func (repo *TractionRepository) GetTractionUsers() []models.TractionUserModel {
	var users []models.TractionUserModel
	result := models.TractionUserModel{}

	iter := repo.C.Find(bson.M{}).Iter()

	for iter.Next(&result) {
		users = append(users, result)
	}

	return users
}

// GetAllTasks all tasks
func (repo *TractionRepository) GetAllTasks() []models.TractionTaskModel {
	var tasks []models.TractionTaskModel
	result := models.TractionTaskModel{}

	iter := repo.C.Find(bson.M{}).Iter()

	for iter.Next(&result) {
		tasks = append(tasks, result)
	}

	return tasks
}

// GetAllMittings tr all mittings
func (repo *TractionRepository) GetAllMittings() []models.TractionMittingModel {
	var mittings []models.TractionMittingModel
	result := models.TractionMittingModel{}

	iter := repo.C.Find(bson.M{}).Iter()

	for iter.Next(&result) {
		mittings = append(mittings, result)
	}

	return mittings
}

// SaveTask saving ti local dataBase
func (repo *TractionRepository) SaveTask(task models.TractionTaskModel) error {

	err := repo.C.Insert(&task)

	return err
}

// func saveToAsana(slice []models.TractionTaskModel) {

// 	fmt.Println("SAVING DATA TO ASANA ITEMS LEFT => " + strconv.Itoa(len(slice)))
// 	if len(slice) == 0 {
// 		return
// 	}
// 	services.AsanaCreateTask(slice[0].DetailsURL, slice[0].Name)
// 	slice = common.RemoveTractionTaskFromArray(slice, 0)
// 	saveToAsana(slice)
// }

// func (repo *TractionRepository) AddUsers(user *models.TractionUserModel) error {
// 	err := repo.C.Insert(user)

// 	return err
// }
