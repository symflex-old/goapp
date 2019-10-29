package main

import (
	"app/src/app"
	"app/src/config"
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
)

const configPath string = "/etc/app.yml"

var appDirectory string

func init()  {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	appDirectory = dir
}

type User struct {
	tableName struct{} `pg:"user"`
	Id	  uuid.UUID
	Email string
	Password string
	Name  string

}

func main()  {

	/*
	config := new(ConfigYaml)
	config.load(appDirectory+configPath)*/

	config := new(config.ConfigYaml)
	config.Load(appDirectory+configPath)
	/*
	db := pg.Connect(&pg.Options{
		User: "postgres",
		Password: "root",
		Database: "docker",
	})
	defer db.Close()

	user1 := &User{
		Name:   "admin",
		Email: "test@test.ru",
		Password: "test",
	}
	err1 := db.Insert(user1)
	if err1 != nil {
		panic(err1)
	}*/


	appService := new(app.AppService)
	appService.Bootstrap(config)
	appService.Process()



	/*

	appService := new(AppService)
	appService.Bootstrap(config)
	err := appService.Process()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error when starting or running http server: %v", err))
	}*/
}
