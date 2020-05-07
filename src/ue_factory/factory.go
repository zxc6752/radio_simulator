package ue_factory

import (
	"fmt"
	"io/ioutil"
	"log"
	"radio_simulator/src/simulator_context"

	"gopkg.in/yaml.v2"
)

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		log.Panic(err.Error())
	}
}

func InitUeContextFactory(f string) *simulator_context.UeContext {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	context := simulator_context.NewUeContext()

	err = yaml.Unmarshal([]byte(content), context)
	checkErr(err)
	return context
}
