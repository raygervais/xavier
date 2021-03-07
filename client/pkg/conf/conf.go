package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration allows us to define a json structure which can be marshalled / un-marshalled with ease.
type Configuration struct {
	ServerLocation string `json:"server"`
}

// ApplicationConfiguration is Singleton instance of Configuration to be passed around server
var ApplicationConfiguration = Configuration{
	ServerLocation: "https://jsonplaceholder.typicode.com/todos/",
}

const (
	// ApplicationName defines the directory name which we'll store all assets to the client application.
	ApplicationName = "/xavier"
	// ConfigurationName defines the name of the config file which will house the Configuration struct in JSON format.
	ConfigurationName = "/conf.json"
)

// DetermineStorageLocation returns the results of UserConfigDir,
// which allows us to support per-os configurations being in the correct location.
func DetermineStorageLocation() (string, error) {
	return os.UserConfigDir()
}

// InitializeConfigurationLocation is a wrapper around the configuration folder handling
// During next application start up, it parses any changes which are made to conf.json
func InitializeConfigurationLocation(path string) error {
	appPath := path + ApplicationName
	if err := createConfigurationFolder(appPath); err != nil {
		return err
	}

	confPath := appPath + ConfigurationName
	if err := createConfigurationFile(confPath, ApplicationConfiguration); err != nil {
		return err
	}

	configuration, err := parseConfigurationFile(confPath)
	if err != nil {
		return err
	}

	ApplicationConfiguration = configuration

	return nil
}

func createConfigurationFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0777); err != nil {
			return fmt.Errorf("Failed to create configuration folder: %s", err)
		}
	}

	return nil
}

func createConfigurationFile(path string, defConf Configuration) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		byteValue, err := json.MarshalIndent(defConf, "", " ")
		if err != nil {
			return err
		}

		_, err = file.Write(byteValue)

		if err != nil {
			return err
		}
	}

	return nil
}

func parseConfigurationFile(path string) (Configuration, error) {
	var configuration Configuration

	file, err := os.Open(path)
	if err != nil {
		return Configuration{}, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &configuration)

	return configuration, nil
}
