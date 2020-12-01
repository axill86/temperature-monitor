package config

import (
	"time"
)

type Config struct {
	Mqtt       MqttConfig
	DeviceName string
	Interval   time.Duration
	I2C        I2CConfig
}

type MqttConfig struct {
	BrokerUrl string
	Topic     string
	DeviceId  string
	Username  string
	Password  string
	RootCa    string
}

type I2CConfig struct {
	Bus     int
	Address uint8
}
