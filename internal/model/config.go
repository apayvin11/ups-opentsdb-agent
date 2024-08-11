package model

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	OpentsdbAddr    string        `toml:"opentsdb_addr"`
	UpsAddr         string        `toml:"ups_addr"`
	UpsTagName      string        `toml:"ups_tag_name"`
	PollingInterval time.Duration `toml:"polling_interval"` // sec
}

func (conf *Config) Validate() error {
	return validation.ValidateStruct(
		conf,
		validation.Field(&conf.OpentsdbAddr, validation.Required),
		validation.Field(&conf.UpsAddr, validation.Required),
		validation.Field(&conf.UpsTagName, validation.Required),
		validation.Field(&conf.PollingInterval, validation.Required, validation.Min(time.Second*5)),
	)
}

func NewConfig(configPath string) (*Config, error) {
	conf := &Config{}
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		return nil, fmt.Errorf("toml decode file config error: %v", err)
	}

	conf.PollingInterval *= time.Second

	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return conf, nil
}
