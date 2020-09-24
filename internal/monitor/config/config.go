package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type Config struct {
	Mqtt       MqttConfig
	DeviceName string
	Interval   time.Duration
}

type MqttConfig struct {
	BrokerUrl string
	Topic     string
	DeviceId  string
	Username  string
	Password  string
	RootCa    string
}

func CreateMqttClient(config *MqttConfig) (mqtt.Client, error) {
	certBytes := []byte(config.RootCa)
	block, _ := pem.Decode(certBytes)
	//load certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(cert)
	mqttClientOpts := mqtt.NewClientOptions()
	mqttClientOpts.SetTLSConfig(&tls.Config{RootCAs: certPool})
	mqttClientOpts.SetPassword(config.Password)
	mqttClientOpts.SetUsername(config.Username)
	mqttClientOpts.AddBroker(config.BrokerUrl)
	return mqtt.NewClient(mqttClientOpts), nil
}
