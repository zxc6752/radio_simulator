package factory

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var SimConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		log.Panic(err.Error())
	}
}

func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	SimConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &SimConfig)
	checkErr(err)
}
