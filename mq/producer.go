package mq

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"rabbit/rabbitmq_example/config"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

func initChannel() bool {
	var err error
	// determine whether this channel is exist
	if channel != nil {
		return true
	}
	conn, err = amqp.Dial(config.RabbitUrl)
	if err != nil {
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		return false
	}
	return true
}

func Publish(exchange, routineKey string, msg []byte) error {
	var (
		payload amqp.Publishing
		err     error
	)
	if !initChannel() {
		return errors.New("init channel failed, err: " + err.Error())
	}
	payload = amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}

	err = channel.Publish(exchange, routineKey, false, false, payload)
	if err != nil {
		return err
	}
	log.Println("publish data")
	return nil
}

