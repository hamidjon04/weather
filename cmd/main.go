package main

import (
	"fmt"
	"weather/api"
	"weather/pkg"
	"weather/service"
	"weather/storage"
	"weather/storage/postgres"
)

func main(){
	log := pkg.InitLogger()
	db, err := postgres.Connect()
	if err != nil{
		log.Error(fmt.Sprintf("Postgres is not connecting: %v", err))
	}
	defer db.Close()

	storage := storage.NewStorage(db, log)
	service := service.NewService(storage, log)
	router := api.Router(service, log)
	if err = router.Run(":8080"); err != nil{
		log.Error(fmt.Sprintf("Error is run router: %v", err))
	}
}