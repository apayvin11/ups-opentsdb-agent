package main

import (
	"fmt"
	"log"
	"time"
	"ups-agent/model"

	"github.com/alex11prog/opentsdb-client/opentsdb"
	"github.com/goburrow/modbus"
)

const (
	db_host        = "127.0.0.1"
	db_port        = 4242
	db_dialTimeout = 3

	upsAddr = "localhost:1502"
)

var opentsdbClient *opentsdb.Client
var modbusClient modbus.Client

func main() {
	opentsdbClient = opentsdb.NewClient(db_host, db_port, db_dialTimeout)
	defer opentsdbClient.Close()

	modbusTcphandler := modbus.NewTCPClientHandler(upsAddr)
	if err := modbusTcphandler.Connect(); err != nil {
		log.Fatal(err)
	}
	defer modbusTcphandler.Close()
	modbusClient = modbus.NewClient(modbusTcphandler)

	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		if res, err := sendUpsParams(readUpsParams()); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%#v\n", res)
		}
		if res, err := sendAlarms(readAlarms()); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%#v\n", res)
		}
	}
}

func readUpsParams() *model.UpsParams {
	var upsParams model.UpsParams
	res, err := modbusClient.ReadHoldingRegisters(0, 70)
	if err != nil {
		log.Fatal(err)
	}
	if err := upsParams.FillFromBytes(res); err != nil {
		log.Fatal(err)
	}
	return &upsParams
}

func getDefaultTags() map[string]interface{} {
	return map[string]interface{}{
		"entity": "localhost_ups_1",
	}
}

func sendUpsParams(p *model.UpsParams) (*opentsdb.PutResponse, error) {
	timestamp := time.Now().UnixMilli()

	metrics := []*opentsdb.UniMetric{
		{
			MetricName: "InputAcVoltage",
			TimeStamp:  timestamp,
			Value:      float64(p.InputAcVoltage),
			Tags:       getDefaultTags(),
		},
		{
			MetricName: "InputAcCurrent",
			TimeStamp:  timestamp,
			Value:      float64(p.InputAcCurrent),
			Tags:       getDefaultTags(),
		},
		{
			MetricName: "BatGroupVoltage",
			TimeStamp:  timestamp,
			Value:      float64(p.BatGroupVoltage),
			Tags:       getDefaultTags(),
		},
		{
			MetricName: "BatGroupCurrent",
			TimeStamp:  timestamp,
			Value:      float64(p.BatGroupCurrent),
			Tags:       getDefaultTags(),
		},
	}
	for i := range p.Batteries {
		batMetrics := []*opentsdb.UniMetric{
			{
				MetricName: fmt.Sprintf("Battery_%d_Voltage", i),
				TimeStamp:  timestamp,
				Value:      float64(p.Batteries[i].Voltage),
				Tags:       getDefaultTags(),
			},
			{
				MetricName: fmt.Sprintf("Battery_%d_Temp", i),
				TimeStamp:  timestamp,
				Value:      float64(p.Batteries[i].Temp),
				Tags:       getDefaultTags(),
			},
			{
				MetricName: fmt.Sprintf("Battery_%d_Resist", i),
				TimeStamp:  timestamp,
				Value:      float64(p.Batteries[i].Resist),
				Tags:       getDefaultTags(),
			},
		}
		metrics = append(metrics, batMetrics...)
	}

	return opentsdbClient.Put(metrics)
}

func readAlarms() *model.Alarms {
	var alarms model.Alarms
	res, err := modbusClient.ReadCoils(0, 3)
	if err != nil {
		log.Fatal(err)
	}
	if err := alarms.FillFromBytes(res); err != nil {
		log.Fatal(err)
	}
	return &alarms
}

func sendAlarms(al *model.Alarms) (*opentsdb.PutResponse, error) {
	timestamp := time.Now().UnixMilli()

	metrics := []*opentsdb.UniMetric{
		{
			MetricName: "Alarm_UpcInBatteryMode",
			TimeStamp:  timestamp,
			Value:      bool2float64(al.UpcInBatteryMode),
			Tags:       getDefaultTags(),
		},
		{
			MetricName: "Alarm_LowBattery",
			TimeStamp:  timestamp,
			Value:      bool2float64(al.LowBattery),
			Tags:       getDefaultTags(),
		},
		{
			MetricName: "Alarm_Overload",
			TimeStamp:  timestamp,
			Value:      bool2float64(al.Overload),
			Tags:       getDefaultTags(),
		},
	}
	return opentsdbClient.Put(metrics)
}

func bool2float64(val bool) float64 {
	if val {
		return 1
	}
	return 0
}
