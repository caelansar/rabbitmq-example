package mq

import (
	"log"
	"time"
)

var done chan bool

func StartConsumer(queue, name string, callback func(msg []byte) bool) {
	log.Println("start consuming")
	msgs, err := channel.Consume(
		queue, name,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err.Error())
		return
	}


	done = make(chan bool)

	//go func() {
	//	for msg := range msgs {
	//		ok := callback(msg.Body)
	//		if !ok {
	//			log.Println("deal message failed")
	//			msg.Nack(false, false)
	//		} else {
	//			msg.Nack(false, false)
	//			log.Println("deal message success")
	//		}
	//
	//	}
	//}()
	go func() {
		for {
			select {
			case <-done:
				log.Println("close channel")
				channel.Close()
				return
			case msg := <-msgs:
				log.Println("msg: ", string(msg.Body))
				ok := callback(msg.Body)
				if !ok {
					log.Println("deal message failed")
					msg.Nack(false, false)
				} else {
					msg.Nack(false, false)
					log.Println("deal message success")
				}
			default:
				time.Sleep(2 * time.Second)
				log.Println("waiting...")

			}
		}
	}()

	//<-done
	//log.Println("close channel")
	//channel.Close()
}

func StopConsumer()  {
	log.Println("stop consuming")
	done <- true
}