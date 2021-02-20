package api

import (
	"fmt"
	"net/http"

	"github.com/raygervais/xavier/client/pkg/conf"
)

type API struct {
	config conf.Configuration
}

func Init(config conf.Configuration) API {
	return API{
		config: config,
	}
}

func (api API) HealthCheck() error {
	_, err := http.Get(api.config.ServerLocation)
	return err
}

func (api API) GetAll() (string, error) {
	resp, err := http.Get(api.config.ServerLocation)
	if err != nil {
		return "", fmt.Errorf("error while reaching out to server %s: %s", api.config.ServerLocation, err)
	}

	defer resp.Body.Close()

	return "", nil
}
