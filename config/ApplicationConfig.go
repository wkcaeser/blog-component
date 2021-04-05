package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ApplicationConf struct {
	Mysql struct {
		Url          string `yaml:"url"`
		MaxOpenConns int    `yaml:"max-open-conns"`
		MaxIdleConns int    `yaml:"max-idle-conns"`
	}
	Cors struct {
		AccessControlAllowOrigin      string `yaml:"access-control-allow-origins"`
		AccessControlAllowMethods     string `yaml:"access-control-allow-methods"`
		AccessControlAllowHeaders     string `yaml:"access-control-allow-headers"`
		AccessControlExposeHeaders    string `yaml:"access-control-expose-headers"`
		AccessControlAllowCredentials bool   `yaml:"access-control-allow-credentials"`
	}
}

var GlobalConf ApplicationConf

func init() {
	configFile, err := ioutil.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}
	log.Println(string(configFile))

	GlobalConf = ApplicationConf{}

	err = yaml.Unmarshal(configFile, &GlobalConf)
	if err != nil {
		panic(err)
	}
}
