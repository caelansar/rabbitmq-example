package config

const (
	RabbitUrl      = "amqp://guest:guest@127.0.0.1:5672/"
	ExchangeName   = "golang"
	TestQueueName  = "golang.test"
	ErrQueueName   = "golang.err"
	TestRoutingKey = "test"
	ErrRoutingKey  = "err"
)
