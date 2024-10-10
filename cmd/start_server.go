package cmd

import (
	"net/http"

	"github.com/Israel-Ferreira/shopping-api/internal/routes"
)

func StartServer(port string) error {
	routes.InitRoutes()

	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}
