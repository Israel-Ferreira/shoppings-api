package main

import (
	"log"

	"github.com/Israel-Ferreira/shopping-api/cmd"
	"github.com/Israel-Ferreira/shopping-api/internal/config"
	"github.com/Israel-Ferreira/shopping-api/internal/databases"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Servidor iniciado na Porta 4000!!")

	godotenv.Load()

	if err := config.LoadEnv(); err != nil {
		log.Fatalln(err)
	}

	databases.CreateMongoConn(config.DbHost, config.DbPort, config.DbUsername, config.DbPassword)

	port := ":4000"

	if err := cmd.StartServer(port); err != nil {
		log.Fatalln(err)
	}
}
