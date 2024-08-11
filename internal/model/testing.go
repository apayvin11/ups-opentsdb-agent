package model

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestUpsParams(t *testing.T) UpsParams {
	return UpsParams{
		InputAcVoltage:  220,
		InputAcCurrent:  2.5,
		BatGroupVoltage: 48,
		BatGroupCurrent: 10,
		Batteries: [NumberOfBatteries]BatteryParams{
			{
				Voltage: 12,
				Temp:    22,
				Resist:  5,
			},
			{
				Voltage: 12.5,
				Temp:    23,
				Resist:  6,
			},
			{
				Voltage: 11.5,
				Temp:    21.5,
				Resist:  5.5,
			},
			{
				Voltage: 12.02,
				Temp:    21,
				Resist:  4.5,
			},
		},
	}
}

func TestUpsParamsToBytes(t *testing.T, params UpsParams) []byte {
	res := make([]byte, 140)
	binary.BigEndian.PutUint32(res[RegInputAcVoltage*2:], math.Float32bits(params.InputAcVoltage))
	binary.BigEndian.PutUint32(res[RegInputAcCurrent*2:], math.Float32bits(params.InputAcCurrent))
	binary.BigEndian.PutUint32(res[RegBatteryGroupVoltage*2:], math.Float32bits(params.BatGroupVoltage))
	binary.BigEndian.PutUint32(res[RegBatteryGroupCurrent*2:], math.Float32bits(params.BatGroupCurrent))

	binary.BigEndian.PutUint32(res[RegBattery1Voltage*2:], math.Float32bits(params.Batteries[0].Voltage))
	binary.BigEndian.PutUint32(res[RegBattery1Temp*2:], math.Float32bits(params.Batteries[0].Temp))
	binary.BigEndian.PutUint32(res[RegBattery1InternalResistance*2:], math.Float32bits(params.Batteries[0].Resist))

	binary.BigEndian.PutUint32(res[RegBattery2Voltage*2:], math.Float32bits(params.Batteries[1].Voltage))
	binary.BigEndian.PutUint32(res[RegBattery2Temp*2:], math.Float32bits(params.Batteries[1].Temp))
	binary.BigEndian.PutUint32(res[RegBattery2InternalResistance*2:], math.Float32bits(params.Batteries[1].Resist))

	binary.BigEndian.PutUint32(res[RegBattery3Voltage*2:], math.Float32bits(params.Batteries[2].Voltage))
	binary.BigEndian.PutUint32(res[RegBattery3Temp*2:], math.Float32bits(params.Batteries[2].Temp))
	binary.BigEndian.PutUint32(res[RegBattery3InternalResistance*2:], math.Float32bits(params.Batteries[2].Resist))

	binary.BigEndian.PutUint32(res[RegBattery4Voltage*2:], math.Float32bits(params.Batteries[3].Voltage))
	binary.BigEndian.PutUint32(res[RegBattery4Temp*2:], math.Float32bits(params.Batteries[3].Temp))
	binary.BigEndian.PutUint32(res[RegBattery4InternalResistance*2:], math.Float32bits(params.Batteries[3].Resist))
	return res
}
