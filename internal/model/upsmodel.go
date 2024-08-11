package model

import (
	"encoding/binary"
	"fmt"
	"math"
)

const NumberOfBatteries = 4

type Alarms struct {
	UpcInBatteryMode bool
	LowBattery       bool
	Overload         bool
}

func GetAlarmsFromBytes(data []byte) (alarms Alarms, err error) {
	if len(data) != 1 {
		err = fmt.Errorf("invalid data len: %d, expected 1", len(data))
		return
	}
	b := data[0]
	alarms.UpcInBatteryMode = b&0b1 != 0
	alarms.LowBattery = b&0b10 != 0
	alarms.Overload = b&0b100 != 0
	return
}

type UpsParams struct {
	InputAcVoltage  float32 // V
	InputAcCurrent  float32 // Amp
	BatGroupVoltage float32 // V
	BatGroupCurrent float32 // Amp
	Batteries       [NumberOfBatteries]BatteryParams
}

type BatteryParams struct {
	Voltage float32
	Temp    float32
	Resist  float32
}

func GetUpsParamsFromBytes(data []byte) (params UpsParams, err error) {
	if len(data) != RegUpsParamsCnt*2 {
		err = fmt.Errorf("invalid data len: %d, expected 140", len(data))
		return
	}
	params.InputAcVoltage = math.Float32frombits(binary.BigEndian.Uint32(data[RegInputAcVoltage*2:]))
	params.InputAcCurrent = math.Float32frombits(binary.BigEndian.Uint32(data[RegInputAcCurrent*2:]))
	params.BatGroupVoltage = math.Float32frombits(binary.BigEndian.Uint32(data[RegBatteryGroupVoltage*2:]))
	params.BatGroupCurrent = math.Float32frombits(binary.BigEndian.Uint32(data[RegBatteryGroupCurrent*2:]))

	for i := range params.Batteries {
		startReg := (RegBattery1Voltage + 0x10*uint16(i)) * 2
		params.Batteries[i].Voltage = math.Float32frombits(binary.BigEndian.Uint32(data[startReg:]))
		params.Batteries[i].Temp = math.Float32frombits(binary.BigEndian.Uint32(data[startReg+4:]))
		params.Batteries[i].Resist = math.Float32frombits(binary.BigEndian.Uint32(data[startReg+8:]))
	}
	return
}
