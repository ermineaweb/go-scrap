package messaging

type Consumer interface {
	Consume(func([]byte))
}

type Producer interface {
	Produce([]byte)
}
