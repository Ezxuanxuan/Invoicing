package tools

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"db"`
	User     string `yaml:"user"`
}

func GetConf(confPath string) *Conf {
	c := new(Conf)
	yamlFile, err := ioutil.ReadFile(confPath)
	fmt.Println(yamlFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
