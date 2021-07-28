package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	ServerHost  string `yaml:"serverhost"`
	ConnectType string `yaml:"connecttype"`
	ServerPort  string `yaml:"serverport"`
}

type ConfigureHelper struct{}

func (this *ConfigureHelper) GetConf() (c *Conf) {
	yamlFile, err := ioutil.ReadFile("../conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(yamlFile))
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
