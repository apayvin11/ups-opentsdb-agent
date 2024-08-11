package agent

import (
	"testing"
	"time"
	"ups-agent/internal/agent/mock/mockmodbus"
	"ups-agent/internal/agent/mock/mocktsdbclient"
)

func TestAgent(t *testing.T, tsdbClient *mocktsdbclient.MockTsdbClient, 
	modbusClient *mockmodbus.MockModbus) *Agent {
	return &Agent{
		tsdbClient: tsdbClient,
		modbusClient: modbusClient,
		upsTagName: "upsTest",
		pollingInterval: time.Minute,
	}
}
