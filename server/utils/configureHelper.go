package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServiceConf struct {
	ServerHost string `yaml:"serverhost"`
	ServerPort string `yaml:"serverport"`
	ListenType string `yaml:"listentype"`
	RedisHost  string `yaml:"redishost"`
	RedisPort  string `yaml:"redisport"`
}

type ConfigureHelper struct{}

func (this *ConfigureHelper) GetConf() (c *ServiceConf) {
	yamlFile, err := ioutil.ReadFile("../conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
