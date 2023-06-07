package api

type ConcurrencyChannel struct {
	msgs   *chan ConcurrencyOperationResult
	cancel *chan bool
}

type ConcurrencyOperationResult struct {
	Result any
	Done   bool
	Error  error
}

func NewConcurrencyChannel() *ConcurrencyChannel {
	ch := &ConcurrencyChannel{}
	msgs := make(chan ConcurrencyOperationResult, 1)
	ch.msgs = &msgs
	cancelled := make(chan bool, 1)
	ch.cancel = &cancelled
	return ch
}

func (c *ConcurrencyChannel) GetConsumer() *Consumer {
	return &Consumer{
		msgs:   c.msgs,
		cancel: c.cancel,
	}
}

func (c *ConcurrencyChannel) GetProducer() *Producer {
	return &Producer{
		msgs: c.msgs,
	}
}

type Producer struct {
	msgs *chan ConcurrencyOperationResult
}

type Consumer struct {
	msgs   *chan ConcurrencyOperationResult
	cancel *chan bool
}

type ConcurrencyOperationCallback func(result ConcurrencyOperationResult)

func (c *Consumer) Consume(callback ConcurrencyOperationCallback) ConcurrencyOperationResult {
	for {
		msg := <-*c.msgs
		callback(msg)
		if msg.Done || msg.Error != nil {
			return msg
		}
	}
}

func (c *Consumer) Cancel() {
	*c.cancel <- true // signal when done
}

func (p *Producer) Produce(result ConcurrencyOperationResult) {
	*p.msgs <- result
}
