package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://test.mosquitto.org:1883"
	topic := "test_topic"
	qos := 0

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		text := fmt.Sprintf("Message sent at: %v", rand.Intn(1000))
		token := client.Publish(topic, byte(qos), false, text)
		token.Wait()
		fmt.Printf("Published: %s\n", text)
		time.Sleep(1 * time.Second)
	}
}
