package agent

import (
	"testing"
	"time"
	"ups-agent/internal/agent/mock/mockmodbus"
	"ups-agent/internal/agent/mock/mocktsdbclient"
	"ups-agent/internal/model"

	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_readAndSendData(t *testing.T) {
	mockTsdbClient := mocktsdbclient.New()
	mockModbus := mockmodbus.New()
	agent := TestAgent(t, mockTsdbClient, mockModbus)

	mockModbus.WriteSingleCoil(model.RegUpsInBatteryMode, 0b00000101)

	upsParams := model.TestUpsParams(t)

	mockModbus.WriteMultipleRegisters(
		model.RegInputAcVoltage,
		model.UpsParamsCnt,
		model.TestUpsParamsToBytes(t, upsParams),
	)

	require.NoError(t, agent.readAndSendData())

	timestamp := time.Now().UnixMilli()
	expected := [][]client.DataPoint{
		{
			{
				Metric:    metricInputAcVoltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.InputAcVoltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricInputAcCurrent,
				Timestamp: timestamp,
				Value:     float64(upsParams.InputAcCurrent),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBatGroupVoltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.BatGroupVoltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBatGroupCurrent,
				Timestamp: timestamp,
				Value:     float64(upsParams.BatGroupCurrent),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_1_Voltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[0].Voltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_1_Temp,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[0].Temp),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_1_Resist,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[0].Resist),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_2_Voltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[1].Voltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_2_Temp,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[1].Temp),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_2_Resist,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[1].Resist),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_3_Voltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[2].Voltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_3_Temp,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[2].Temp),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_3_Resist,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[2].Resist),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_4_Voltage,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[3].Voltage),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_4_Temp,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[3].Temp),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricBattery_4_Resist,
				Timestamp: timestamp,
				Value:     float64(upsParams.Batteries[3].Resist),
				Tags:      agent.getDefaultTags(),
			},
		},
		{
			{
				Metric:    metricAlarmUpcInBatteryMode,
				Timestamp: timestamp,
				Value:     float64(1),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricAlarmLowBattery,
				Timestamp: timestamp,
				Value:     float64(0),
				Tags:      agent.getDefaultTags(),
			},
			{
				Metric:    metricAlarmOverload,
				Timestamp: timestamp,
				Value:     float64(1),
				Tags:      agent.getDefaultTags(),
			},
		},
	}

	assert.Equal(t, expected, mockTsdbClient.GetData())
}
