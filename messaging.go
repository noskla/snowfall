package main

import (
	"github.com/streadway/amqp"
	"log"
)

var RabbitMQInstance *amqp.Connection
var RabbitMQChannel  *amqp.Channel

func InitRabbitMQ() bool {

	var err error
	var rabbitMQProtocol string = "amqp://"
	if Cfg.RabbitMQSecure {
		rabbitMQProtocol = "amqps://"
	}

	rabbitMQUrl := rabbitMQProtocol + Cfg.RabbitMQUsername + ":" + Cfg.RabbitMQPassword + "@" + Cfg.RabbitMQAddress + ":" + Cfg.RabbitMQPort + "/"

	RabbitMQInstance, err = amqp.Dial(rabbitMQUrl)
	if errorOccurred(err, false) {
		log.Println("Warning: RabbitMQ could not be contacted. Discord integration will be broken.")
		return false
	}

	RabbitMQChannel, err = RabbitMQInstance.Channel()
	if errorOccurred(err, false) {
		log.Println("Warning: RabbitMQ channel could not be opened. Discord integration will be broken.")
		return false
	}

	RabbitMQChannel.QueueDeclare("SendDiscordValidationMessage", false, false, false, false, nil)
	//RabbitMQChannel.QueueDeclare("DiscordValidationConfirmation", false, false, false, false, nil)
	//discordValConfirm, err := RabbitMQChannel.Consume("DiscordValidationConfirmation", "", true, false, false, false, nil)
	//if errorOccurred(err, false) {
	//	log.Println("Warning: Could not consume RabbitMQ channel. Discord integration will be broken.")
	//	return false
	//}

	//go ConfirmDiscordRoutine(discordValConfirm)

	return true

}
