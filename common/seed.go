package common

import (
	"asanaSync/models"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// SeedAsanaAndTracUsers creates all necessary data for the app
func SeedAsanaAndTracUsers() {
	dbSession := GetDBSession()

	defer dbSession.Close()
	//TractionUserModel
	asanaUserCol := dbSession.DB(AppConfig.DataBaseName).C(AppConfig.CollUserAsana)
	asanaProjectCol := dbSession.DB(AppConfig.DataBaseName).C(AppConfig.CollProjectAsana)
	tractionCol := dbSession.DB(AppConfig.DataBaseName).C(AppConfig.CollUserTraction)
	//tractionTaskCol := dbSession.DB(AppConfig.DataBaseName).C(AppConfig.CollTaskTraction)
	tractionMittingCol := dbSession.DB(AppConfig.DataBaseName).C(AppConfig.CollMittingTraction)

	asanaUserCol.RemoveAll(bson.M{})
	tractionCol.RemoveAll(bson.M{})
	asanaProjectCol.RemoveAll(bson.M{})
	tractionMittingCol.RemoveAll(bson.M{})
	//tractionTaskCol.RemoveAll(bson.M{})

	err := asanaUserCol.Insert(
		&models.AsanaUser{
			ID:           1126471228603132,
			Gid:          "1126471228603132",
			Email:        "alex@vetintegrations.com",
			Name:         "Alex Balabanov",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1126472345209962,
			Gid:          "1126472345209962",
			Email:        "bogdan@vetintegrations.com",
			Name:         "Bogdan",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1126472345980362,
			Gid:          "1126472345980362",
			Email:        "sergei@vetintegrations.com",
			Name:         "Sergei Nechaev",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1126472539603771,
			Gid:          "1126472539603771",
			Email:        "ivanzak@vetintegrations.com",
			Name:         "Ivan Zakharenkov",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1127754072456244,
			Gid:          "1127754072456244",
			Email:        "bill@vetintegrations.com",
			Name:         "Bill Griffin",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1127799994857801,
			Gid:          "1127799994857801",
			Email:        "noemail@vetintegrations.com",
			Name:         "Private User",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           235992055555881,
			Gid:          "235992055555881",
			Email:        "olegborzykh@gmail.com",
			Name:         "oleg borzykh",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1132152485916575,
			Gid:          "1132152485916575",
			Email:        "nancy@vetintegrations.com",
			Name:         "Nancy Stewart",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1132152485916590,
			Gid:          "1132152485916590",
			Email:        "e_nantz@hotmail.com",
			Name:         "Emmitt Nantz",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1137905023577368,
			Gid:          "1137905023577368",
			Email:        "jeremy@vetintegrations.com",
			Name:         "Jeremy",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1138033300878896,
			Gid:          "1138033300878896",
			Email:        "nastya@vetintegrations.com",
			Name:         "Nastya",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1142155389074767,
			Gid:          "1142155389074767",
			Email:        "malinovskiynewpost@gmail.com",
			Name:         "malinovskiynewpost",
			ResourceType: "user",
		},
		&models.AsanaUser{
			ID:           1142373512896311,
			Gid:          "1142373512896311",
			Email:        "support@vetintegrations.comm",
			Name:         "Support",
			ResourceType: "user",
		},
	)

	if err != nil {
		fmt.Println("error in seed seedAsanaUsers")
		log.Fatal(err)
	}

	err = tractionCol.Insert(&models.TractionUserModel{
		UserID: 306503,
		Name:   "Alex Balabanov",
		Email:  "alex@vetintegrations.com",
	},
		&models.TractionUserModel{
			UserID: 308831,
			Name:   "Bogdan Babich",
			Email:  "bogdan@vetintegrations.com",
		},
		&models.TractionUserModel{
			UserID: 306505,
			Name:   "Sergei Nechaev",
			Email:  "sergei@vetintegrations.com",
		},
		&models.TractionUserModel{
			UserID: 296468,
			Email:  "ivanzak@vetintegrations.com",
			Name:   "Ivan Zak",
		},
		&models.TractionUserModel{
			UserID: 306507,
			Email:  "olegborzykh@gmail.com",
			Name:   "Oleg Borzykh",
		},
		&models.TractionUserModel{
			UserID: 296460,
			Email:  "nancy@vetintegrations.com",
			Name:   "Nancy Stewar",
		},
		&models.TractionUserModel{
			UserID: 296520,
			Email:  "e_nantz@hotmail.com",
			Name:   "Emmitt Nantz",
		},
		&models.TractionUserModel{
			UserID: 306509,
			Email:  "jeremy@vetintegrations.com",
			Name:   "Jeremy Halperin",
		},
		&models.TractionUserModel{
			UserID: 308833,
			Email:  "nastya@vetintegrations.com",
			Name:   "Nastya Babich",
		},
		&models.TractionUserModel{
			UserID: 324950,
			Email:  "malinovskiynewpost@gmail.com",
			Name:   "Vladimir Malinovskiy",
		},
	)

	if err != nil {
		fmt.Println("error in seed SeedAsanaAndTracUsers")
		log.Fatal(err)
	}

	err = asanaProjectCol.Insert(&models.AsanaProject{
		ID:           1131118892044634,
		Gid:          "1131118892044634",
		Name:         "Marketing L10 To Do",
		ResourceType: "project",
	},
		&models.AsanaProject{
			ID:           1131118414232707,
			Gid:          "1131118414232707",
			Name:         "VIS Platform L10 To Do",
			ResourceType: "project",
		},
		&models.AsanaProject{
			ID:           1131118892044651,
			Gid:          "1131118892044651",
			Name:         "Support L10 To Do",
			ResourceType: "project",
		},
		&models.AsanaProject{
			ID:           1131118892044658,
			Gid:          "1131118892044658",
			Name:         "R&D To Do",
			ResourceType: "project",
		},
		&models.AsanaProject{
			ID:           1142153483755225,
			Gid:          "1142153483755225",
			Name:         "VIS L10",
			ResourceType: "project",
		},
	)

	if err != nil {
		fmt.Println("error in seed SeedASANA Projects")
		log.Fatal(err)
	}

	err = tractionMittingCol.Insert(&models.TractionMittingModel{
		ID:   35535,
		Name: "VIS Platform L10 To Do",
	},
		&models.TractionMittingModel{
			ID:   35543,
			Name: "Marketing L10 To Do",
		},
		&models.TractionMittingModel{
			ID:   35544,
			Name: "Support L10 To Do",
		},
		&models.TractionMittingModel{
			ID:   35545,
			Name: "R&D To Do",
		},
		&models.TractionMittingModel{
			ID:   35536,
			Name: "VIS L10",
		},
	)

	if err != nil {
		fmt.Println("error in seed SeedTraction Mittings")
		log.Fatal(err)
	}

	fmt.Println("Seed completed!")
}
