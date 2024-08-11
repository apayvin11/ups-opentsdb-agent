package main

import (
	"fmt"
	"log"
	"ups-agent/internal/agent"
	"ups-agent/internal/model"

	"github.com/goburrow/modbus"
)

const confPath = "conf/config.toml"

func main() {
	conf, err := model.NewConfig(confPath)
	if err != nil {
		log.Fatal(err)
	}

	modbusTcphandler := modbus.NewTCPClientHandler(conf.UpsAddr)
	if err := modbusTcphandler.Connect(); err != nil {
		log.Fatal(err)
	}
	defer modbusTcphandler.Close()
	modbusClient := modbus.NewClient(modbusTcphandler)

	agent, err := agent.New(conf, modbusClient)
	if err != nil {
		log.Fatal(err)
	}

	agent.Start()

	fmt.Println("press enter to quit")
	var s string
	fmt.Scanln(&s)
}
