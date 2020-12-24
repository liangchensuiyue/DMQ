package main

import (
	"client_shell/mycmd"
	"client_shell/mylog"
	pd "client_shell/proto"
	"context"
	"fmt"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"

	"google.golang.org/grpc"
)

var ttype string
var topic string
var group string
var address string

func main() {

	var err error
	cmd := mycmd.NewCmd()
	topic, err = cmd.Get("topic")
	if err != nil {
		mylog.Error("没有指定目标话题!")
		os.Exit(0)
	}
	ttype, _ := cmd.Get("type")
	if ttype != "consumer" && ttype != "producer" {
		mylog.Error("未选择一个客户端类型(consumer/producer)")
	}
	address, _ = cmd.Get("address")
	group, err = cmd.Get("group")
	if err != nil {
		mylog.Error("没有指定消费者群组!")
		os.Exit(0)
	}
	if ttype == "consumer" {
		startConsume()
	} else {
		startProduce()
	}
}
func startProduce() {
	mylog.Info(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		mylog.Error(fmt.Sprintf("Connect success(%s)", address))
		os.Exit(1)
	} else {
		mylog.Success(fmt.Sprintf("Connect success(%s)", address))
	}
	//网络延迟关闭
	defer conn.Close()

	//获得grpc句柄
	c := pd.NewDMQFollowerServiceClient(conn)

	//通过句柄调用函数
	res, err := c.ClientYieldMsgDataRequest(context.Background())
	if err != nil {
		mylog.Error(fmt.Sprintf("Connect transport(%s) ", address) + err.Error())
	}
	for {
		var message string
		fmt.Scanln(&message)
		if message == "exit" {
			res.CloseAndRecv()
			os.Exit(0)
		}
		str := strings.TrimSpace(message)
		if len(str) == 0 {
			continue
		}
		err = res.Send(&pd.MessageData{Topic: topic, Message: str})
		if err != nil {
			os.Exit(1)
		}
	}
}
func startConsume() {
	var nodeid string
	u1 := uuid.NewV4()
	str, _ := u1.MarshalText()
	nodeid = string(str)
	mylog.Info(address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		mylog.Error(fmt.Sprintf("Connect success(%s)", address))
		os.Exit(1)
	} else {
		mylog.Success(fmt.Sprintf("Connect success(%s)", address))
	}
	//网络延迟关闭
	defer conn.Close()

	//获得grpc句柄
	c := pd.NewDMQFollowerServiceClient(conn)

	//通过句柄调用函数
	res, err := c.ClientConsumeData(context.Background(), &pd.ClientRegistToFollower{Nodeid: nodeid, Groupname: group, Topic: topic})
	if err != nil {
		mylog.Error(fmt.Sprintf("Connect transport(%s) ", address) + err.Error())
		os.Exit(1)
	}
	for info, _ := res.Recv(); info != nil; info, _ = res.Recv() {
		if info.Errno == 0 {
			fmt.Printf("(%s) %s\n", nodeid, info.Data.Message)
		}
	}
}
