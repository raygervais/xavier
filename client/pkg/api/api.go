package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/raygervais/xavier/client/pkg/conf"
)

// API is a wrapper instance which contains application configuration
type API struct {
	config conf.Configuration
}

// Init creates the API instance configured
// to user requirements, ex. Server URL
func Init(config conf.Configuration) API {
	return API{
		config: config,
	}
}

// HealthCheck verifies that the CLI can connect
// to the server specified in the configuration
func (api API) HealthCheck() error {
	resp, err := http.Get(api.config.ServerLocation)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	if resp.StatusCode != 200 {
		return fmt.Errorf(
			"Received different status code than 200. %s\n%s",
			body,
			err,
		)
	}

	println(body)

	return nil
}
