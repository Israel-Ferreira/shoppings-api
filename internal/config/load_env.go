package config

import (
	"errors"
	"os"
)

var DbHost, DbPort, DbUsername, DbPassword string

func LoadEnv() error {
	DbHost = os.Getenv("MONGODB_HOST")

	if DbHost == "" {
		return errors.New("erro: O host do banco de dados não foi informado")
	}

	DbPassword = os.Getenv("MONGODB_PASSWORD")

	if DbPassword == "" {
		return errors.New("erro: A senha do banco de dados não foi informada")
	}

	DbPort = os.Getenv("MONGODB_PORT")

	if DbPort == "" {
		DbPort = "27017"
	}

	DbUsername = os.Getenv("MONGODB_USERNAME")

	if DbUsername == "" {
		return errors.New("erro: O usuário do banco de dados não foi informado")
	}

	return nil

}
