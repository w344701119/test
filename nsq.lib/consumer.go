package nsq_lib

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

//callback
type Callback func(content string) bool

//ConsumerHandle
type ConsumerHandle struct {
	cb Callback
}

//处理消息
func (h *ConsumerHandle) HandleMessage(msg *nsq.Message) error {
	h.cb(string(msg.Body))
	return nil
}

//Consumer
type Consumer struct {
	HandleFunc Callback
	Topic      string
	Channel    string
	Address    []string
}

//初始化消费者
func (c *Consumer) Init() error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	consumer, err := nsq.NewConsumer(c.Topic, c.Channel, cfg)
	if err != nil {
		return err
	}
	//屏蔽系统日志
	consumer.SetLogger(nil, 0)
	//添加消费者接口
	handle := &ConsumerHandle{
		cb: c.HandleFunc,
	}
	consumer.AddHandler(handle)
	//建立NSQ连接
	if err := consumer.ConnectToNSQDs(c.Address); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) Run() {
	for range time.NewTicker(time.Second * 60).C {
		fmt.Println(222)
	}
}
