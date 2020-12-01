/*
notification package contains components required to send mqtt notifications to broker
*/
package notification

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// Data contains notification data
type Data struct {
	Temperature float32
	Humidity    float32
	Timestamp   time.Time
}

type Sender interface {
	Notify(data Data) error
}

func createMqttClient(sslCert []byte, username, password, brokerUrl string) (mqtt.Client, error) {
	block, _ := pem.Decode(sslCert)
	//load certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(cert)
	mqttClientOpts := mqtt.NewClientOptions()
	mqttClientOpts.SetTLSConfig(&tls.Config{RootCAs: certPool})
	mqttClientOpts.SetPassword(password)
	mqttClientOpts.SetUsername(username)
	mqttClientOpts.AddBroker(brokerUrl)
	return mqtt.NewClient(mqttClientOpts), nil
}

func NewMqttSender(sslCert []byte, username, password, brokerUrl, topic string) (*SenderImpl, error) {
	mqttClient, err := createMqttClient(sslCert, username, password, brokerUrl)
	if err != nil {
		return nil, err
	}
	return &SenderImpl{client: mqttClient, topic: topic}, nil
}

type SenderImpl struct {
	client mqtt.Client
	topic  string
}

func (sender *SenderImpl) Notify(data Data) error {
	//Open connection
	if token := sender.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if token := sender.client.Publish(sender.topic, 0, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
