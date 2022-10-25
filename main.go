package main

import (
	db2 "amani/db"
	"amani/db_manager"
	"amani/handler"
)

func main() {
	db := db2.SetUpDb("project.db")
	dm := db_manager.NewDbManager(db)

	h := handler.NewHandler(dm)
	h.Start()
}
