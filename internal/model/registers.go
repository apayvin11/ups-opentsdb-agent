package model

const (
	// holding registers
	RegInputAcVoltage      uint16 = 0x0000
	RegInputAcCurrent      uint16 = 0x0002
	RegBatteryGroupVoltage uint16 = 0x0004
	RegBatteryGroupCurrent uint16 = 0x0006

	RegBattery1Voltage            uint16 = 0x0010
	RegBattery1Temp               uint16 = 0x0012
	RegBattery1InternalResistance uint16 = 0x0014

	RegBattery2Voltage            uint16 = 0x0020
	RegBattery2Temp               uint16 = 0x0022
	RegBattery2InternalResistance uint16 = 0x0024

	RegBattery3Voltage           uint16 = 0x0030
	RegBattery3Temp               uint16 = 0x0032
	RegBattery3InternalResistance uint16 = 0x0034

	RegBattery4Voltage            uint16 = 0x0040
	RegBattery4Temp               uint16 = 0x0042
	RegBattery4InternalResistance uint16 = 0x0044

	RegUpsParamsCnt = 0x0046
	UpsParamsCnt = 16
)

const(
	// discrete inputs
	RegUpsInBatteryMode uint16 = 0x0000
	RegLowBattery       uint16 = 0x0001
	RegOverload         uint16 = 0x0002

	RegAlarmsCnt = 0x0003
)
