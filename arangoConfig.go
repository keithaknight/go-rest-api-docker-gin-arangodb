package main

import (
	"os"
	"strings"
)

type arangoDbConfig struct {
	endpointUrls []string
	userName     string
	password     string
	databaseName string
}

func newArangoDbConfig() arangoDbConfig {
	envUrls := os.Getenv("ARANGODB_URLS")
	var config arangoDbConfig

	if envUrls != "" {
		urls := strings.Split(envUrls, ",")

		config = arangoDbConfig{
			endpointUrls: urls,
			userName:     os.Getenv("ARANGODB_USER"),
			password:     os.Getenv("ARANGODB_PASSWORD"),
			databaseName: os.Getenv("ARANGODB_DATABASE"),
		}
	} else {
		config = arangoDbConfig{
			endpointUrls: []string{""}, //"https://db.domain.com:8529"
			userName:     "",
			password:     "",
			databaseName: "",
		}
	}

	return config
}
