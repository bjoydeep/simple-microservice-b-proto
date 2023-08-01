package transport

import (
	"encoding/json"
	"fmt"

	"github.com/bjoydeep/simple-microservice-b-proto/pkg/config"
	"github.com/bjoydeep/simple-microservice-b-proto/pkg/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(client mqtt.Client, topic string, messageChan chan mqtt.Message) {
	//defer wg.Done()

	//helpful:
	//func (mqtt.Client).Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
	//Subscribe starts a new subscription. Provide a MessageHandler to be executed when a message
	//is published on the topic provided, or nil for the default handler.
	//If options.OrderMatters is true (the default) then callback must not block or call functions
	//within this package that may block (e.g. Publish) other than in a new go routine.
	//callback must be safe for concurrent use by multiple goroutines
	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		messageChan <- msg
		go processMessages(client, messageChan)
	})
	token.Wait()
	println("Subscribed to topic successfully: ", topic)
}

func Publish(client mqtt.Client, jsonData []byte, topic string) {

	//println("Publishing messages..-----", string(jsonData))
	//helpful: https://github.com/eclipse/paho.mqtt.golang/blob/master/client.go#L767-L776
	token := client.Publish(topic, 1, true, jsonData)
	//call blocks till the message is sent to the broker
	token.Wait()
	println("Published messages", string(jsonData), " to topic: ", topic)

}

func processMessages(client mqtt.Client, messageChan <-chan mqtt.Message) {
	var user model.User
	for msg := range messageChan {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		// Message processing here
		err := json.Unmarshal(msg.Payload(), &user)
		if err != nil {
			println("Error marshalling JSON data: ", err.Error())
		} else {
			fmt.Println("Printing User as recieved: ", user)
			//sets the user status to processed
			user = model.SetStatus(user, "Processed")
			//uses that updated user to publish to the topic
			jsonBytes, err := json.Marshal(user)
			if err != nil {
				println("Error marshalling JSON data: ", err.Error())
			}
			Publish(client, jsonBytes, config.Cfg.BrokerPubTopic)
		}
	}
}
