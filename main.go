package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbit/rabbitmq_example/config"
	"rabbit/rabbitmq_example/mq"
	"time"
)

func HandleMsg(msg []byte) bool {
	// only print message body for test
	var msgData mq.MsgData
	err := json.Unmarshal(msg, msgData)
	if err != nil {
		log.Println(err.Error())
		err := mq.Publish(config.ExchangeName, config.ErrRoutingKey, []byte(
			fmt.Sprintf("deal message failed, message is %v, err is %v", string(msg), err.Error())),
		)
		if err != nil {
			log.Println("send message to err queue failed, err:", err.Error())
			return false
		}
		log.Println("send message to error queue")
		return false
	}
	fmt.Println(msgData)
	return true
}

func main() {
	msg1 := mq.MsgData{
		Id:  1,
		Str: "message1",
	}
	msg2 := mq.MsgData{
		Id:  2,
		Str: "message2",
	}

	msgs := []mq.MsgData{msg1, msg2}

	go func() {
		for _, msg := range msgs {
			msgByte, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
				return
			}
			err = mq.Publish(config.ExchangeName, config.TestRoutingKey, msgByte)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)

	mq.StartConsumer(
		config.TestQueueName,
		"test_consumer",
		HandleMsg,
	)

	time.Sleep(5 * time.Second)

	mq.StopConsumer()
}
