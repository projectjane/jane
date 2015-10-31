package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadConfig(location string) (config Config) {

	file, err := ioutil.ReadFile(location)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
	}

	config.BambooChannels = make(map[string]string)
	config.BambooChannels["*"] = "#devops"
	config.BambooChannels["HTML"] = "#random"

	return config

}

func CheckConfig(location string) (exists bool) {
	exists = true
	if _, err := os.Stat(location); os.IsNotExist(err) {
		exists = false
	}
	return exists
}
