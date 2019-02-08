package main

import (
	"log"

	"./db"
)

func main() {
	users := []db.User{
		db.User{Name: "Steve", Mobile: "+222222"},
		db.User{Name: "Mark", Mobile: "+111111"},
		db.User{Name: "Jimmy", Mobile: "+33333"},
	}

	groups := []db.Group{
		db.Group{Name: "Group1"},
		db.Group{Name: "Group2"},
		db.Group{Name: "Group3"},
	}

	alerts := []db.Alert{
		db.Alert{ID: "ID1"},
		db.Alert{ID: "ID2"},
		db.Alert{ID: "ID3"},
	}

	err := db.AddUsers(users)
	if err != nil {
		log.Printf("Error adding users: %+v", err)
	} else {
		log.Println("Users added successfully")
	}

	err = db.AddGroups(groups)
	if err != nil {
		log.Printf("Error adding groups: %+v", err)
	} else {
		log.Println("Groups added successfully")
	}

	err = db.AddAlerts(alerts)
	if err != nil {
		log.Printf("Error adding alerts: %+v", err)
	} else {
		log.Println("Alerts added successfully")
	}

}
