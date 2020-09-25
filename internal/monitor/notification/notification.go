/*
notification package contains components required to send mqtt notifications to broker
*/
package notification

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// Data contains notification data
type Data struct {
	Temperature float32
	Timestamp   time.Time
}

type Sender interface {
	Notify(data Data) error
}

func NewMqttSender(client mqtt.Client, topic string) *SenderImpl {
	return &SenderImpl{client: client, topic: topic}
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
