package api

import (
	"fmt"
	"io/ioutil"
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
		return "", fmt.Errorf(
			"error while reaching out to server %s: %s",
			api.config.ServerLocation,
			err,
		)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf(
			"error while attempting to read response body: %s",
			err,
		)
	}

	return string(body), nil
}
