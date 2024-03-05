package main

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://test.mosquitto.org:1883"
	topic := "test_topic"
	qos := 0

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("subscriber")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	token := client.Subscribe(topic, byte(qos), func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Block indefinitely until interrupted
	select {}
}
