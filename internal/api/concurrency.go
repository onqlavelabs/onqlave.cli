package api

type ConcurrencyOperationResult struct {
	Result any
	Done   bool
	Error  error
}

type ConcurrencyChannel struct {
	msg    *chan ConcurrencyOperationResult
	cancel *chan bool
}

func NewConcurrencyChannel() *ConcurrencyChannel {
	ch := &ConcurrencyChannel{}
	msg := make(chan ConcurrencyOperationResult, 1)
	ch.msg = &msg
	cancelled := make(chan bool, 1)
	ch.cancel = &cancelled
	return ch
}

func (c *ConcurrencyChannel) GetConsumer() *Consumer {
	return &Consumer{
		msg:    c.msg,
		cancel: c.cancel,
	}
}

func (c *ConcurrencyChannel) GetProducer() *Producer {
	return &Producer{
		msg: c.msg,
	}
}

type Producer struct {
	msg *chan ConcurrencyOperationResult
}

type Consumer struct {
	msg    *chan ConcurrencyOperationResult
	cancel *chan bool
}

type ConcurrencyOperationCallback func(result ConcurrencyOperationResult)

func (c *Consumer) Consume(callback ConcurrencyOperationCallback) ConcurrencyOperationResult {
	for {
		msg := <-*c.msg
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
	*p.msg <- result
}
