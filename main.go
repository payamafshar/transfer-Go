package main

import (
	"ReservApp/src/api"
	"ReservApp/src/cmd"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("invalid env file")
	}
	err = cmd.SetupAppConfig()
	config := cmd.GetAppConfig()
	if err != nil {
		panic("Failed to setup app config")
	}
	fmt.Println("Application run  sda on", config.Api.ApiPort)
	fmt.Println("hello")
	err = api.SetupServer(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to run server", err))
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to run server", err))
	}

}
