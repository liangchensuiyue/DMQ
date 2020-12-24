package service

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"server/mylog"
	cachepd "server/service/proto/cache"
	pd "server/service/proto/follower"
	headerpd "server/service/proto/header"
	"server/utils"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type clientNode struct {
	nodeId    string
	groupname string
	topic     string
	out       pd.DMQFollowerService_ClientConsumeDataServer
}

var cacheServiceClient cachepd.CacheCenterServiceClient
var followerToHeaderRequestClient headerpd.DMQHeaderService_FollowerToHeaderRequestDataRequestClient
var topicList map[string]map[string][]clientNode = make(map[string]map[string][]clientNode)
var followerMsgChan chan *headerpd.Response = make(chan *headerpd.Response, 100)
var headerService headerpd.DMQHeaderServiceClient
var followerconfig *utils.MyConfig
var clientNodeDead map[string]bool = make(map[string]bool, 0)
var SendLock *sync.Mutex = &sync.Mutex{}

// 客户端生产数据
func (this *server) ClientYieldMsgDataRequest(in pd.DMQFollowerService_ClientYieldMsgDataRequestServer) error {
	res, err := headerService.FollowerYieldMsgDataRequest(context.Background())
	if err != nil {
		in.SendAndClose(&pd.Response{Errno: 1, Errmsg: err.Error()})
		return nil
	}
	for info, _ := in.Recv(); info != nil; info, _ = in.Recv() {
		// 将客户端的生产的数据转发给 header
		SendLock.Lock()
		err1 := res.Send(&headerpd.MessageData{Topic: info.Topic, Message: info.Message})
		SendLock.Unlock()
		if err1 != nil {
			in.SendAndClose(&pd.Response{Errno: 1, Errmsg: err.Error()})
			return nil
		}
		// 发给 header 节点
	}
	in.SendAndClose(&pd.Response{Errno: 2, Errmsg: "生产通道已关闭已关闭"})
	return nil
}

// 客户端消费数据
func (this *server) ClientConsumeData(request *pd.ClientRegistToFollower, in pd.DMQFollowerService_ClientConsumeDataServer) error {
	/*
			    string groupname = 1;       // 消费者组名
		    string topic     = 2;       // 话题名称
		    string address = 3;         // 绑定的ip 地址
		    int32 port    = 4;         // 绑定的 端口
		    int64 offset = 5;
	*/
	var topic string
	var group string
	topic = request.Topic
	group = request.Groupname
	if topicList[topic] == nil {
		topicList[topic] = make(map[string][]clientNode, 0)
	}
	if topicList[topic][group] == nil {
		topicList[topic][group] = make([]clientNode, 0)
	}
	topicList[topic][group] = append(topicList[topic][group], clientNode{topic: topic, groupname: group, out: in, nodeId: request.Nodeid})
	res, err := cacheServiceClient.Get(context.Background(), &cachepd.Request{Key: request.Topic + "_" + request.Groupname})
	if err != nil {
		return nil
	}
	_v, _ := strconv.Atoi(res.Data)
	offset := int64(_v)
	fmt.Println(offset, "offset")
	followerToHeaderRequestClient.Send(&headerpd.FollowerToHeaderRequestData{Groupname: group, Topic: topic, Offset: offset})
	for {
		time.Sleep(time.Second * 1)
		_, ok := clientNodeDead[request.Nodeid]
		if ok {
			delete(clientNodeDead, request.Nodeid)
			return nil
		}
	}
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// time.Sleep(time.Second * 10)
	return nil
}

// 客户端关闭链接
func (this *server) ClientCloseChannel(ctx context.Context, request *pd.ClientRegistToFollower) (*pd.Response, error) {
	// for info, _ := in.Recv(); info != nil; info, _ = in.Recv() {

	// 	// 将数据添加到管道
	// 	msgChan <- *info
	// 	// 管道
	// }
	// in.SendAndClose(&pd.Response{Errno: 2, Errmsg: "生产通道已关闭已关闭"})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// time.Sleep(time.Second * 4)
	return &pd.Response{}, nil
}
func linkHeader(conf *utils.MyConfig) {
	//客户端连接服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.G_Header_Address, conf.G_Header_Port), grpc.WithInsecure())
	if err != nil {
		mylog.Error("网络异常: " + err.Error())
	}
	//网络延迟关闭
	// defer conn.Close()

	c := headerpd.NewDMQHeaderServiceClient(conn)
	headerService = c
	res, err := c.FollowerToHeaderRequestDataRequest(context.Background())
	if err != nil {
		mylog.Error("链接失败: " + err.Error())
		os.Exit(1)
	}
	res.Send(&headerpd.FollowerToHeaderRequestData{Address: conf.G_Bind_Address, Port: int32(conf.G_Bind_Port)})
	info, e := res.Recv()
	if e != nil {
		mylog.Error(fmt.Sprintf("Connect error(%s:%d): %s", conf.G_Bind_Address, conf.G_Bind_Port, e.Error()))
		os.Exit(1)
	}
	if info.Errno != 0 {
		mylog.Error(fmt.Sprintf("Connect error(%s:%d): %s", conf.G_Bind_Address, conf.G_Bind_Port, info.Errmsg))
		os.Exit(1)
	}
	mylog.Success(fmt.Sprintf("Connect success(%s:%d): %s", conf.G_Bind_Address, conf.G_Bind_Port, info.Errmsg))
	followerToHeaderRequestClient = res
}
func linkCache(conf *utils.MyConfig) {
	// 链接缓存中心
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.G_Cache_Address, conf.G_Cache_Port), grpc.WithInsecure())
	if err != nil {
		mylog.Error("网络异常: " + err.Error())
	}

	c := cachepd.NewCacheCenterServiceClient(conn)
	cacheServiceClient = c
}
func receiveHeaderData() {
	for info, _ := followerToHeaderRequestClient.Recv(); info != nil; info, _ = followerToHeaderRequestClient.Recv() {
		// if info.Errno == 0 {
		followerMsgChan <- info
		// }2
	}
}
func acceptData() {
	var response *headerpd.Response
	var data *headerpd.MessageData
	for {
		response = <-followerMsgChan
		data = response.Data
		fmt.Println("receive: ", data)
		if response.Errno != 0 {
			consumerList := topicList[data.Topic]
			if consumerList == nil {
				continue
			}
			for p, _ := range consumerList {
				if consumerList[p] != nil {
					for i := 0; i < len(consumerList[p]); i++ {
						consumerList[p][i].out.Send(&pd.Response{Errno: 1, Errmsg: response.Errmsg})
						clientNodeDead[consumerList[p][i].nodeId] = true
					}
				}
			}
			delete(topicList, data.Topic)
			continue
		}
		// fmt.Println(*data)
		for {
			consumerList := topicList[data.Topic]
			if consumerList == nil {
				break
			}
			if data.Des != "" {
				fmt.Println(data)
				if consumerList[data.Des] == nil || len(consumerList[data.Des]) == 0 {
					break
				}
				i := rand.Intn(len(consumerList[data.Des]))
				err := consumerList[data.Des][i].out.Send(&pd.Response{Errno: 0, Data: &pd.MessageData{Topic: data.Topic, Message: data.Message, Length: data.Length - 1}})

				// 发送成功后 更新 offset
				if err == nil {
					/*
						res, _err := cacheServiceClient.Get(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, data.Des)})
						if _err != nil {
							mylog.Error("cache transport error: " + _err.Error())
							break
						}
						_num, _ := strconv.Atoi(res.Data)
						value := int64(_num) + data.Length
						cacheServiceClient.Put(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, data.Des), Value: fmt.Sprintf("%d", value)})
					*/

					cacheServiceClient.Add(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, data.Des), Value: fmt.Sprintf("%d", data.Length)})
				} else {
					// 发送失败(该链接断开)
					if len(consumerList[data.Des]) == 1 {
						delete(topicList[data.Topic], data.Des)
						if len(topicList[data.Topic]) == 0 {
							delete(topicList, data.Topic)

							// 向 header 发送注销该topic的请求
							headerService.FollowerCancelTopicRequest(context.Background(), &headerpd.FollowerCancelTopic{Address: followerconfig.G_Bind_Address, Port: int32(followerconfig.G_Bind_Port), Topic: data.Topic})
						}
					} else {
						if len(consumerList[data.Des])-1 == i {
							topicList[data.Topic][data.Des] = topicList[data.Topic][data.Des][:len(consumerList[data.Des])-1]
						} else {
							topicList[data.Topic][data.Des] = append(topicList[data.Topic][data.Des][:i], topicList[data.Topic][data.Des][i+1:]...)
						}
					}
				}
				break
			}

			consumerList = topicList[data.Topic]
			for p, _ := range consumerList {
				if consumerList[p] == nil || len(consumerList[p]) == 0 {
					continue
				}

				for {
					consumerList = topicList[data.Topic]
					if consumerList[p] == nil || len(consumerList[p]) <= 0 {
						break
					}
					i := rand.Intn(len(consumerList[p]))

					err := consumerList[p][i].out.Send(&pd.Response{Errno: 0, Data: &pd.MessageData{Topic: data.Topic, Message: data.Message, Length: data.Length - 1}})

					// 发送成功后 更新 offset
					if err == nil {
						/*
							res, _err := cacheServiceClient.Get(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, p)})
							if _err != nil {
								mylog.Error("cache transport error: " + _err.Error())
								continue
							}
							_num, _ := strconv.Atoi(res.Data)
							value := int64(_num) + data.Length
							fmt.Println(_num, value, data.Length, res)
							cacheServiceClient.Put(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, p), Value: fmt.Sprintf("%d", value)})
						*/
						cacheServiceClient.Add(context.Background(), &cachepd.Request{Key: fmt.Sprintf("%s_%s", data.Topic, p), Value: fmt.Sprintf("%d", data.Length)})

						break
					} else {
						fmt.Println("发送失败 => ", consumerList[p][i].nodeId)
						// 发送失败(该链接断开)
						if len(consumerList[p]) == 1 {
							delete(topicList[data.Topic], p)
							if len(topicList[data.Topic]) == 0 {
								delete(topicList, data.Topic)
								// 向 header 发送注销该topic的请求
								headerService.FollowerCancelTopicRequest(context.Background(), &headerpd.FollowerCancelTopic{Address: followerconfig.G_Bind_Address, Port: int32(followerconfig.G_Bind_Port), Topic: data.Topic})
							}
						} else {
							if len(consumerList[p])-1 == i {
								topicList[data.Topic][p] = topicList[data.Topic][p][:len(consumerList[p])-1]
							} else {
								topicList[data.Topic][p] = append(topicList[data.Topic][p][:i], topicList[data.Topic][p][i+1:]...)
							}
						}
					}
				}

			}
			break
		}
	}
}
func StartFollower(conf *utils.MyConfig) {
	followerconfig = conf
	// initGlobalVar()
	// go channelAroused()
	linkHeader(conf)
	go receiveHeaderData()
	go acceptData()
	linkCache(conf)

	//创建网络
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.G_Bind_Address, conf.G_Bind_Port))
	if err != nil {
		fmt.Println("网络错误", err)
	}

	//创建grpc的服务
	srv := grpc.NewServer()

	//注册服务
	pd.RegisterDMQFollowerServiceServer(srv, &server{})
	mylog.Info(fmt.Sprintf("listen: %s:%d\n", conf.G_Bind_Address, conf.G_Bind_Port))
	//等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		mylog.Error("启动失败: " + err.Error())
	}
	ln.Close()
	mylog.Info("ending")

}
