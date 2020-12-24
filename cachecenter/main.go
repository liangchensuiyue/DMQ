package main

import (
	"bytes"
	pd "cachecenter/proto"
	"context"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

var lock *sync.Mutex = &sync.Mutex{}
var datadir = "./data.txt"

var cache map[string]string

func (this *server) Put(ctx context.Context, in *pd.Request) (out *pd.Response, err error) {
	cache[in.Key] = in.Value
	fmt.Println("put: ", in.Key, in.Value)
	return &pd.Response{Errno: 0}, nil
}
func (this *server) Add(ctx context.Context, in *pd.Request) (out *pd.Response, err error) {
	lock.Lock()
	defer lock.Unlock()
	_, ok := cache[in.Key]
	if !ok {
		cache[in.Key] = "0"
	}
	value, err := strconv.Atoi(in.Value)
	ovalue, _ := strconv.Atoi(cache[in.Key])
	if err != nil {
		return &pd.Response{Errno: 1}, nil
	}
	cache[in.Key] = fmt.Sprintf("%d", ovalue+value)
	fmt.Println("add: ", in.Key, in.Value)
	return &pd.Response{Errno: 0}, nil
}
func (this *server) Get(ctx context.Context, in *pd.Request) (out *pd.Response, err error) {
	value, ok := cache[in.Key]
	fmt.Println("get: ", in.Key, in.Value)
	if !ok {
		cache[in.Key] = "0"
		return &pd.Response{Errno: 1}, nil
	}
	return &pd.Response{Errno: 0, Data: value}, nil
}
func loadFile() {
	_, err := os.Stat(datadir)
	if os.IsNotExist(err) {
		cache = make(map[string]string, 0)
		return
	}
	// 读取钱包
	content, err := ioutil.ReadFile(datadir)
	if err != nil {
		log.Panic(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader(content))

	err = decoder.Decode(&cache)
	if err != nil {
		log.Panic(err)
	}
}
func saveToFile() {
	/*
		如果 Encode/Decode 类型是interface或者struct中某些字段是interface{}的时候
		需要在gob中注册interface可能的所有实现或者可能类型
	*/
	var content bytes.Buffer

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(&cache)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(datadir, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
func RegularTimeTask() {

	// 5 秒更新一次数据到dish
	time.Sleep(time.Second * 5)
	saveToFile()
	RegularTimeTask()
}

func main() {
	//创建网络
	ln, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println("网络错误", err)
	}

	//创建grpc的服务
	srv := grpc.NewServer()

	go RegularTimeTask()
	loadFile() // 加载缓存的数据
	fmt.Println(cache)

	//注册服务
	pd.RegisterCacheCenterServiceServer(srv, &server{})

	//等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}

}
