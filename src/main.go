package main

import (
	"fmt"
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

func main()  {

	config := new(ConfigYaml)
	config.load(appDirectory+configPath)

	appService := new(AppService)
	appService.Bootstrap(config)
	err := appService.Process()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error when starting or running http server: %v", err))
	}
}
