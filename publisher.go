package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://test.mosquitto.org:1883"
	topic := "test_topic"
	qos := 1

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		// text := fmt.Sprintf("Message sent at: %v", rand.Intn(1000))
		// text := `{"id": "u8ia","data": "10","unit": "kg"}`

		text := `{
			"id": "somemessageid101",
			"node_id": "somenodeid404",
			"method": "COMMAND",
			"category": "SENSOR",
			"action": "START",
			"payload": "{\"sensor_id\": \"sens101\",\"data\": \"65\",\"unit\": \"kg\",\"value\":\"ON\"}",
			"timestamp": "Tue Mar 05 2024 14:33:06 GMT+0530 (India Standard Time)"
		}`
		token := client.Publish(topic, byte(qos), false, text)
		token.Wait()
		fmt.Printf("Published: %s\n", text)
		time.Sleep(1 * time.Second)
	}
}
