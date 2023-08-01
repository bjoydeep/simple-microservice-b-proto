package transport

import (
	"fmt"

	"github.com/bjoydeep/simple-microservice-b-proto/pkg/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
// 	//println("Received message kind: ", string(msg.Payload()), " from topic: ", msg.Topic(),
// 	//	", Is duplicate: ", msg.Duplicate(), ", Is retained: ", msg.Retained(), ", QoS setting: ", msg.Qos())
// }

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQ Broker ---")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection to MQBroker lost---: %v", err)
}

var BrokerClient mqtt.Client
var MessageChan chan mqtt.Message

func SetupTransport() {
	var broker = config.Cfg.BrokerHost
	var port = config.Cfg.BrokerPort
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("microservice-b")
	// opts.SetUsername("bar")
	// opts.SetPassword("baz")
	//opts.SetDefaultPublishHandler(messagePubHandler)
	//opts.OrderMatters = true
	opts.CleanSession = false
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//println("Connected to MQ Broker Post: ", client.IsConnected())
	BrokerClient = client

	// Create a channel to receive messages
	MessageChan = make(chan mqtt.Message, 1)

}
