package main

import (
	"encoding/json"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	Id   string
	Data string
	Unit string
}

func main() {
	broker := "tcp://test.mosquitto.org:1883"
	topic := "test_topic"
	qos := 1

	var sensorData SensorData

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("subscriber")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	token := client.Subscribe(topic, byte(qos), func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		json.Unmarshal(msg.Payload(), &sensorData)

		fmt.Println(sensorData.Id, sensorData.Data, sensorData.Unit)
	})
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Block indefinitely until interrupted
	select {}
}
