package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Jeffail/gabs"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	Id   string
	Data string
	Unit string
}

type MqttMessage struct {
	ID        string `json:"id"`
	NODE_ID   string `json:"node_id"`
	METHOD    string `json:"method"`
	CATEGORY  string `json:"category"`
	ACTION    string `json:"action"`
	PAYLOAD   string `json:"payload"`
	TIMESTAMP string `json:"timestamp"`
}

func main() {
	broker := "tcp://test.mosquitto.org:1883"
	topic := "test_topic"
	qos := 1

	var mqttMessage MqttMessage

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("subscriber")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	token := client.Subscribe(topic, byte(qos), func(client MQTT.Client, msg MQTT.Message) {
		// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		json.Unmarshal(msg.Payload(), &mqttMessage)

		fmt.Println(mqttMessage)

		// parsing the payload object (unkown format)
		parcedPayload, err := gabs.ParseJSON([]byte(mqttMessage.PAYLOAD))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(parcedPayload.Path("sensor_id"))
	})
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Block indefinitely until interrupted
	select {}
}
