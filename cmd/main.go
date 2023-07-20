package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bjoydeep/simple-microservice-b-proto/pkg/config"
	"github.com/bjoydeep/simple-microservice-b-proto/pkg/model"
	"github.com/bjoydeep/simple-microservice-b-proto/pkg/transport"
)

func loadConfig() {
	err := config.InitConfig()
	if err != nil {
		println("Error while loading the config: ", err)
	} else {
		println("-------------------------------------")
		println("Configuration Details are : ")
		println("-------------------------------------")
		println("BrokerHost is: ", config.Cfg.BrokerHost)
		println("Topic is: ", config.Cfg.BrokerSubTopic)
		println("Topic is: ", config.Cfg.BrokerPubTopic)
		println("Broker Port is: ", config.Cfg.BrokerPort)

		println("-------------------------------------")
	}
}

func main() {
	// Create a new Gin router
	//router := gin.Default()

	//var wg sync.WaitGroup

	//load the application config details before doing anything else
	loadConfig()
	// Initialize MQTT client
	transport.SetupTransport()

	//Initialized the DB connections
	//storage.SetupStorage()
	//Sets up gorm
	model.SetupModel()

	// Launch goroutines and pass the WaitGroup to each function
	//wg.Add(1)
	go transport.Subscribe(transport.BrokerClient, config.Cfg.BrokerSubTopic, transport.MessageChan)
	go transport.ProcessMessages(transport.BrokerClient, transport.MessageChan)

	// Wait for all goroutines to finish
	//wg.Wait()

	// Create a channel to receive signals (e.g., Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received (e.g., Ctrl+C)
	<-sigChan

	// Define API endpoints - to add versioning
	//router.GET("/users", handler.GetUsers)
	//router.POST("/users", handler.AddUser)
	//router.GET("/user/:id", handler.GetUser)

	// Start the server
	//router.Run(":8080")
}
