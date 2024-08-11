package agent

import (
	"testing"
	"ups-agent/internal/agent/mock/mockmodbus"
	"ups-agent/internal/agent/mock/mocktsdbclient"
	"ups-agent/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_readAndSendData(t *testing.T) {
	mockTsdbClient := mocktsdbclient.New()
	mockModbus := mockmodbus.New()
	agent := TestAgent(t, mockTsdbClient, mockModbus)

	mockModbus.WriteSingleCoil(model.RegUpsInBatteryMode, 0b00000101)

	mockModbus.WriteMultipleRegisters(
		model.RegInputAcVoltage,
		model.UpsParamsCnt,
		model.TestUpsParamsToBytes(t, model.TestUpsParams(t)),
	)

	require.NoError(t, agent.readAndSendData())

	res := mockTsdbClient.GetData()

	assert.Equal(t, 2, len(res))
	assert.Equal(t, model.UpsParamsCnt, len(res[0]))
	assert.Equal(t, model.RegAlarmsCnt, len(res[1]))
}
