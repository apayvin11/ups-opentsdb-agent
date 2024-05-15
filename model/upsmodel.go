package model

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BatteryParams struct {
	Voltage float32 `json:"voltage" example:"12"`
	Temp    float32 `json:"temp" example:"24"`
	Resist  float32 `json:"resist" example:"5"`
}

type Alarms struct {
	UpcInBatteryMode bool `json:"upc_in_battery_mode" example:"false"`
	LowBattery       bool `json:"low_battery" example:"false"`
	Overload         bool `json:"overload" example:"false"`
}

func (a *Alarms) FillFromBytes(data []byte) error {
	if len(data) != 1 {
		return fmt.Errorf("invalid data len: %d, expected 1", len(data))
	}
	b := data[0]
	if b&0b1 > 0 {
		a.UpcInBatteryMode = true
	}
	if b&0b10 > 0 {
		a.UpcInBatteryMode = true
	}
	if b&0b100 > 0 {
		a.UpcInBatteryMode = true
	}
	return nil
}

type UpsParams struct {
	InputAcVoltage  float32          `json:"input_ac_voltage" example:"220"` // V
	InputAcCurrent  float32          `json:"input_ac_current" example:"5"`   // Amp
	BatGroupVoltage float32          `json:"bat_group_voltage" example:"48"` // V
	BatGroupCurrent float32          `json:"bat_group_current" example:"0"`  // Amp
	Batteries       [4]BatteryParams `json:"batteries"`
}

func (up *UpsParams) FillFromBytes(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.BigEndian, &up.InputAcVoltage); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &up.InputAcCurrent); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &up.BatGroupVoltage); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &up.BatGroupCurrent); err != nil {
		return err
	}
	buf.Next(16) // skip empty bytes
	for i := range up.Batteries {
		if err := binary.Read(buf, binary.BigEndian, &up.Batteries[i].Voltage); err != nil {
			return err
		}
		if err := binary.Read(buf, binary.BigEndian, &up.Batteries[i].Temp); err != nil {
			return err
		}
		if err := binary.Read(buf, binary.BigEndian, &up.Batteries[i].Resist); err != nil {
			return err
		}
		buf.Next(20) // skip empty bytes
	}
	return nil
}
