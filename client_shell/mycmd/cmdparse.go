package mycmd

import (
	"errors"
	"flag"
)

var kv map[string]string

func init() {
	kv = make(map[string]string, 0)
}

type Mycmd struct{}

func (cmd *Mycmd) Get(k string) (string, error) {
	value, ok := kv[k]
	if ok {
		return value, nil
	} else {
		return value, errors.New("not found")
	}
}

func NewCmd() *Mycmd {
	var a1 string
	var a2 string
	var a3 string
	var a4 string
	flag.StringVar(&a1, "type", "consumer", "consumper/producer")
	flag.StringVar(&a2, "topic", "", "目标话题")
	flag.StringVar(&a3, "group", "nil", "消费者群组")
	flag.StringVar(&a4, "h", "nil", "远端服务地址(host:port)")
	flag.Parse()
	kv["type"] = a1
	kv["topic"] = a2
	kv["group"] = a3
	kv["address"] = a4
	return &Mycmd{}
}
