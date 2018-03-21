package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type config struct {
	SlkToken string `json:"slackToken"`
}

var (
	aconfig *config
	once    sync.Once
)

func Config() *config {
	once.Do(func() {
		aconfig = &config{}
	})
	return aconfig
}

func (config *config) SlackToken() string {
	return config.SlkToken
}

func (config *config) Load(filename string) bool {
	filestring, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read config")
		return false
	}

	err = json.Unmarshal(filestring, config)
	if err != nil {
		fmt.Println("Unable to parse config")
		return false
	}

	return true
}
