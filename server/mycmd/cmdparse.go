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
	var configfile string
	var cmd string
	flag.StringVar(&configfile, "config", "./config.conf", "配置文件路径")
	flag.StringVar(&cmd, "cmd", "false", "cli 运行")
	flag.Parse()
	kv["configfile"] = configfile
	kv["useCmd"] = cmd
	return &Mycmd{}
}
