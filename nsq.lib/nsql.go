package nsq_lib

import (
	"math/rand"
	"time"

	//matcher "NLU/src/matcher"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"

	//nsqpkg "gopkg/nsq"
	//"gopkg/utils/common"
	"github.com/nsqio/go-nsq"
)

type NsqInfo struct {
	Com         Consumer      // nsq消息消费者
	Producer    *nsq.Producer // nsq生产者
	MasterTopic string
}

var NsqInfoObj = NsqInfo{}

func GetChannelName() string {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(1000000)
	ch := fmt.Sprintf("channel%d%d", rnd, time.Now().Unix())
	return ch
}

func (self *NsqInfo) InitNsq() bool {
	var err error
	nsq_host := beego.AppConfig.DefaultString("nsq_host", "")
	nsq_port := beego.AppConfig.DefaultString("nsq_port", "")
	if nsq_host == "" || nsq_port == "" {
		logs.Error("nsq读取配置文件出错")
		return false
	}
	addr := fmt.Sprintf("%s:%s", nsq_host, nsq_port)
	// 初始化nsq接收
	ch := GetChannelName()
	logs.Info("init nsq, ch:%s", ch)
	var defaultTopic = beego.AppConfig.DefaultString("nsq_topic", "")
	self.Com = Consumer{
		Topic:      defaultTopic,
		Channel:    ch,
		Address:    []string{addr},
		HandleFunc: HandleUpdateInfo,
	}
	err = self.Com.Init()
	if err != nil {
		logs.Error("nsq init err:", err)
		return false
	}
	// 初始化nsq发送
	pdConfig := nsq.NewConfig()

	if self.Producer, err = nsq.NewProducer(addr, pdConfig); err != nil {
		logs.Error("init producer error: %v", err)
		return false
	}
	//默认topic
	self.MasterTopic = defaultTopic

	return true
}

//发送nsql信息
func (self *NsqInfo) SendNsq(data string) bool {
	if self.Producer == nil {
		logs.Info("nsq producer is not init")
		return false
	}
	if err := self.Producer.Publish(self.MasterTopic, []byte(data)); err != nil {
		logs.Error("send nsq err:", err, "topic:", self.MasterTopic, "data:", data)
	} else {
		logs.Info("send nsq success:", "topic:", self.MasterTopic, "data:", data)
	}

	return true
}

func HandleUpdateInfo(content string) bool {
	logs.Info("receive nsq topic:%s, msg:%s", NsqInfoObj.Com.Topic, content)
	js, err := simplejson.NewJson([]byte(content))
	if err != nil {
		logs.Error("json解析错误 : %v", err)
		return false
	}
	body_js := js.Get("body")
	ext_js := body_js.Get("ext")
	action := body_js.Get("action").MustString()
	if action != "update_info" {
		logs.Error("receive invalid action: %s", action)
		return false
	}
	body_js.Set("time", time.Now().String())
	js.Set("ip", NsqInfoObj.Com.Address[0])
	js.Set("topic", NsqInfoObj.MasterTopic)
	ext_js.Set("status", "success")
	var b []byte
	b, err = js.Encode()
	if err != nil {
		logs.Error("encode result js failed, %v", err)
	}
	NsqInfoObj.SendNsq(string(b))
	return true
}
