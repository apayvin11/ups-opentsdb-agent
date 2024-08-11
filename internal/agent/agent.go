package agent

import (
	"fmt"
	"log"
	"time"
	"ups-agent/internal/model"

	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/bluebreezecf/opentsdb-goclient/config"
	"github.com/goburrow/modbus"
)

type Agent struct {
	tsdbClient      client.Client
	modbusClient    modbus.Client
	upsTagName      string
	pollingInterval time.Duration
}

func New(conf *model.Config, modbusClient modbus.Client) (*Agent, error) {
	tsdbClient, err := client.NewClient(config.OpenTSDBConfig{
		OpentsdbHost: conf.OpentsdbAddr,
	})
	if err != nil {
		return nil, err
	}

	if err = tsdbClient.Ping(); err != nil {
		return nil, err
	}
	return &Agent{
		tsdbClient:      tsdbClient,
		modbusClient:    modbusClient,
		upsTagName:      conf.UpsTagName,
		pollingInterval: conf.PollingInterval,
	}, nil
}

// Start starts polling the device and sending data to the opentsdb
func (a *Agent) Start() {
	go func() {
		ticker := time.NewTicker(a.pollingInterval)
		for range ticker.C {
			if err := a.readAndSendData(); err != nil {
				log.Println(err)
			}
		}
	}()
}

func (a *Agent) readAndSendData() error {
	upsParams, err := a.readUpsParams()
	if err != nil {
		return err
	}
	if err := a.sendUpsParams(upsParams); err != nil {
		return err
	}

	alarms, err := a.readAlarms()
	if err != nil {
		return err
	}

	if err := a.sendAlarms(alarms); err != nil {
		return err
	}

	return nil
}

func (a *Agent) readUpsParams() (model.UpsParams, error) {
	res, err := a.modbusClient.ReadHoldingRegisters(model.RegInputAcVoltage, model.RegUpsParamsCnt)
	if err != nil {
		return model.UpsParams{}, err
	}
	return model.GetUpsParamsFromBytes(res)
}

func (a *Agent) sendUpsParams(p model.UpsParams) error {
	timestamp := time.Now().UnixMilli()
	metrics := make([]client.DataPoint, 0, model.UpsParamsCnt)
	metrics = append(metrics, []client.DataPoint{
		{
			Metric:    metricInputAcVoltage,
			Timestamp: timestamp,
			Value:     float64(p.InputAcVoltage),
			Tags:      a.getDefaultTags(),
		},
		{
			Metric:    metricInputAcCurrent,
			Timestamp: timestamp,
			Value:     float64(p.InputAcCurrent),
			Tags:      a.getDefaultTags(),
		},
		{
			Metric:    metricBatGroupVoltage,
			Timestamp: timestamp,
			Value:     float64(p.BatGroupVoltage),
			Tags:      a.getDefaultTags(),
		},
		{
			Metric:    metricBatGroupCurrent,
			Timestamp: timestamp,
			Value:     float64(p.BatGroupCurrent),
			Tags:      a.getDefaultTags(),
		},
	}...)

	for i := range p.Batteries {
		batMetrics := []client.DataPoint{
			{
				Metric:    fmt.Sprintf("Battery_%d_Voltage", i+1),
				Timestamp: timestamp,
				Value:     float64(p.Batteries[i].Voltage),
				Tags:      a.getDefaultTags(),
			},
			{
				Metric:    fmt.Sprintf("Battery_%d_Temp", i+1),
				Timestamp: timestamp,
				Value:     float64(p.Batteries[i].Temp),
				Tags:      a.getDefaultTags(),
			},
			{
				Metric:    fmt.Sprintf("Battery_%d_Resist", i+1),
				Timestamp: timestamp,
				Value:     float64(p.Batteries[i].Resist),
				Tags:      a.getDefaultTags(),
			},
		}
		metrics = append(metrics, batMetrics...)
	}

	_, err := a.tsdbClient.Put(metrics, "details")
	return err

}

func (a *Agent) readAlarms() (model.Alarms, error) {
	res, err := a.modbusClient.ReadCoils(model.RegUpsInBatteryMode, model.RegAlarmsCnt)
	if err != nil {
		return model.Alarms{}, err
	}
	return model.GetAlarmsFromBytes(res)
}

func (a *Agent) sendAlarms(al model.Alarms) error {
	timestamp := time.Now().UnixMilli()

	metrics := []client.DataPoint{
		{
			Metric:    "Alarm_UpcInBatteryMode",
			Timestamp: timestamp,
			Value:     bool2float64(al.UpcInBatteryMode),
			Tags:      a.getDefaultTags(),
		},
		{
			Metric:    "Alarm_LowBattery",
			Timestamp: timestamp,
			Value:     bool2float64(al.LowBattery),
			Tags:      a.getDefaultTags(),
		},
		{
			Metric:    "Alarm_Overload",
			Timestamp: timestamp,
			Value:     bool2float64(al.Overload),
			Tags:      a.getDefaultTags(),
		},
	}
	_, err := a.tsdbClient.Put(metrics, "details")
	return err
}

func (a *Agent) getDefaultTags() map[string]string {
	return map[string]string{
		"entity": a.upsTagName,
	}
}

func bool2float64(val bool) float64 {
	if val {
		return 1.0
	}
	return 0.0
}
