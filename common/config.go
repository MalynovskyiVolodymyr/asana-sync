package common

import (
	"asanaSync/models"
	"fmt"

	"gopkg.in/mgo.v2"
)

type applicationConfiguration struct {
	DataBaseName        string
	CollTaskAsana       string
	CollUserAsana       string
	CollProjectAsana    string
	CollTaskTraction    string
	CollUserTraction    string
	CollMittingTraction string
}

// USER_WAN_NOT_FOND_ERROR error result
const USER_WAN_NOT_FOND_ERROR = "USER_WAN_NOT_FOND_ERROR"

// AppConfig application configuration
var AppConfig applicationConfiguration

func init() {
	AppConfig = applicationConfiguration{
		DataBaseName:        "asanaSyncDataBase",
		CollTaskAsana:       "asanaTasks",
		CollTaskTraction:    "tractionTasks",
		CollUserAsana:       "asanaUsers",
		CollUserTraction:    "tractionUsers",
		CollProjectAsana:    "asanaProjects",
		CollMittingTraction: "tractionMittings",
	}
}

// GetDBSession get database session
func GetDBSession() *mgo.Session {
	session, err := mgo.Dial("localhost") //mgo.Dial("localhost")//MONGO_URL//MONGODB_URI//os.Getenv()

	if err != nil {
		fmt.Println("Error cannot connect to the Database")
	}

	return session
}

// RemoveTractionTaskFromArray traction tools
func RemoveTractionTaskFromArray(slice []models.TractionTaskModel, index int) []models.TractionTaskModel {
	return append(slice[:index], slice[index+1:]...)
}

// PrepareMittingsAndProjects mixing
func PrepareMittingsAndProjects(mittings []models.TractionMittingModel, projects []models.AsanaProject) []models.MittingAndProjectModel {
	var result []models.MittingAndProjectModel
	var temp models.MittingAndProjectModel

	for _, value := range mittings {
		for _, value2 := range projects {
			if value.Name == value2.Name {
				temp = models.MittingAndProjectModel{
					ProjectID: value2.ID,
					MittingID: value.ID,
				}
				result = append(result, temp)
			}
		}
	}

	return result
}
