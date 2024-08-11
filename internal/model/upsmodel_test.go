package model_test

import (
	"testing"
	"ups-agent/internal/model"

	"github.com/stretchr/testify/assert"
)

func Test_GetAlarmsFromBytes(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		isValid  bool
		expected model.Alarms
	}{
		{
			name:    "valid, all alarms",
			data:    []byte{0b00000111},
			isValid: true,
			expected: model.Alarms{
				UpcInBatteryMode: true,
				LowBattery:       true,
				Overload:         true,
			},
		},
		{
			name:    "invalid, empty data",
			isValid: false,
		},
		{
			name:    "valid, alarm UpcInBatteryMode",
			data:    []byte{0b00000001},
			isValid: true,
			expected: model.Alarms{
				UpcInBatteryMode: true,
			},
		},
		{
			name:    "valid, alarm LowBattery",
			data:    []byte{0b00000010},
			isValid: true,
			expected: model.Alarms{
				LowBattery: true,
			},
		},
		{
			name:    "valid, alarm Overload",
			data:    []byte{0b00000100},
			isValid: true,
			expected: model.Alarms{
				Overload: true,
			},
		},
		{
			name:     "valid, all alarms false",
			data:     []byte{0},
			isValid:  true,
			expected: model.Alarms{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			alarms, err := model.GetAlarmsFromBytes(tc.data)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, alarms)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_GetUpsParamsFromBytes(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		isValid  bool
		expected func() model.UpsParams
	}{
		{
			name:    "valid",
			data:    model.TestUpsParamsToBytes(t, model.TestUpsParams(t)),
			isValid: true,
			expected: func() model.UpsParams {
				return model.TestUpsParams(t)
			},
		},
		{
			name:    "invalid data len",
			data:    make([]byte, 15),
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := model.GetUpsParamsFromBytes(tc.data)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected(), params)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
