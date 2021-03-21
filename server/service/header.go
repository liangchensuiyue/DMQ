package service

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"server/mycrypto"
	"server/mylog"
	"server/service/exchange"
	headerpd "server/service/proto/header"
	pd "server/service/proto/header"
	"server/service/tx"
	"server/utils"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

type FollowerNode struct {
	address   string   // follower 节点绑定的ip地址
	port      int      // follower 节点绑定的端口
	topicList []string // foolower 注册的topic
}
type FollowerNodeList []*FollowerNode

type Topic struct {
	file *os.File
	lock *sync.Mutex // 锁
}

var version int = 0
var followerList FollowerNodeList
var topicMap map[string]*Topic
var headerconfig *utils.MyConfig
var PrepareQueue []*tx.Tx = make([]*tx.Tx, 0)

// 记录消费者对应的话题偏移量;隔一定的周期就写入磁盘
// var offsetCache map[string]int
var topicLock map[string]*sync.Mutex

var followerOutputPipe map[string]*pd.DMQHeaderService_FollowerToHeaderRequestDataRequestServer

var msgChan chan pd.MessageData

// 给 follower 节点新增topic
func (fnode *FollowerNode) AddTopic(topic string) {
	for i := 0; i < len(fnode.topicList); i++ {
		if topic == fnode.topicList[i] {
			return
		}
	}
	fnode.topicList = append(fnode.topicList, topic)
}

func (fnode *FollowerNode) HaveTopic(topic string) bool {
	for i := 0; i < len(fnode.topicList); i++ {
		if topic == fnode.topicList[i] {
			return true
		}
	}
	return false
}

// 给 follower 节点注销对应 topic
func (fnode *FollowerNode) RemoveTopic(topic string) {
	for i := 0; i < len(fnode.topicList); i++ {
		if fnode.topicList[i] == topic {
			if i != len(fnode.topicList)-1 {
				fnode.topicList = append(fnode.topicList[:i], fnode.topicList[i+1:]...)
			} else {
				fnode.topicList = fnode.topicList[:i]
			}
			break
		}
	}
}

// 注册 follower 节点
func (flist *FollowerNodeList) Add(address string, port int) error {
	for i := 0; i < len(*flist); i++ {
		if (*flist)[i].address == address && (*flist)[i].port == port {
			return errors.New("该follower节点已被注册!")
		}
	}
	newfollowerNode := FollowerNode{address, port, make([]string, 0)}
	*flist = append(*flist, &newfollowerNode)
	return nil
}

// 注销 follower 节点
func (flist *FollowerNodeList) Remove(n *FollowerNode) error {
	for i := 0; i < len(*flist); i++ {
		if (*flist)[i].address == n.address && (*flist)[i].port == n.port {
			if i != len(*flist)-1 {
				*flist = append((*flist)[:i], (*flist)[i+1:]...)
			} else {
				*flist = (*flist)[:i]
			}
		}
	}
	return nil
}

//一个打招呼的函数

//rpc
//函数关键字（对象）函数名（客户端发送过来的内容 ， 返回给客户端的内容） 错误返回值

//grpc
//函数关键字 （对象）函数名 （cotext，客户端发过来的参数 ）（发送给客户端的参数，错误）
func (this *server) FollowerCancelTopicRequest(ctx context.Context, request *pd.FollowerCancelTopic) (out *pd.Response, err error) {
	for i := 0; i < len(followerList); i++ {
		if followerList[i].address == request.Address && followerList[i].port == int(request.Port) {
			mylog.Info(fmt.Sprintf("follower(%s:%d) topic:%s 已被注销 ", request.Address, request.Port, request.Topic))
			followerList[i].RemoveTopic(request.Topic)
		}
	}
	return &pd.Response{}, nil
}

func (this *server) PingPong(ctx context.Context, request *pd.PingPongData) (out *pd.PingPongData, err error) {
	out = &pd.PingPongData{AliveTime: 10}
	return out, nil
}

func (this *server) CommitTx(ctx context.Context, request *pd.TxData) (out *pd.Response, err error) {
	tx.CommitTx(request.Txid)
	out = &pd.Response{Errno: 0}
	return out, nil
}
func (this *server) Prepare(ctx context.Context, request *pd.TxData) (out *pd.Response, err error) {
	tx.PrepareTx(&tx.Tx{TxId: request.Txid, Topic: request.Topic, Msg: request.Msg})
	out = &pd.Response{Errno: 0}
	return out, nil
}
func EnterQueue(request *pd.MessageData) {
	request.Length = int64(len([]byte(request.Message)) + 1)

	// 将数据写到dish
	// writeDataToTopic(info.Topic, info.Message)
	currentTxid := tx.GetCurrentTxId()
	txdata := &tx.Tx{}
	if currentTxid == "" {
		txdata.TxId = fmt.Sprintf("%d-%d-%d", 0, time.Now().Unix(), 1)
	} else {
		if len(PrepareQueue) == 0 {
			_arr := strings.Split(currentTxid, "-")
			code, _ := strconv.Atoi(_arr[2])
			txdata.TxId = fmt.Sprintf("%d-%d-%d", version, time.Now().Unix(), code+1)
		} else {
			_arr := strings.Split(PrepareQueue[len(PrepareQueue)-1].TxId, "-")
			code, _ := strconv.Atoi(_arr[2])
			txdata.TxId = fmt.Sprintf("%d-%d-%d", version, time.Now().Unix(), code+1)
		}

	}
	txdata.Topic = request.Topic
	txdata.Msg = request.Message
	PrepareQueue = append(PrepareQueue, txdata)
}
func (this *server) Transfer2Master(ctx context.Context, request *pd.MessageData) (out *pd.Response, err error) {
	EnterQueue(request)
	out = &pd.Response{Errno: 0}
	return out, nil
}
func (this *server) GetTxDataByTx(ctx context.Context, request *pd.TxData) (out *pd.TxDatas, err error) {
	txs, _ := tx.GetTxDatasByTxId(request.Txid)
	out = new(pd.TxDatas)
	for _, v := range txs {
		out.Txs = append(out.Txs, &pd.TxData{
			Txid:  v.TxId,
			Topic: v.Topic,
			Msg:   v.Msg,
		})
	}
	return out, nil
}
func (this *server) EnterCluster(ctx context.Context, request *pd.HeaderInfo) (out *pd.HeaderInfo, err error) {
	nodes := exchange.GetHeaders()
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Address == request.Address && nodes[i].Port == request.Port {
			return &pd.HeaderInfo{}, nil
		}
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", request.Address, request.Port), grpc.WithInsecure())
	if err != nil {
		mylog.Error(fmt.Sprintf("%s:%d  %s", request.Address, request.Port, "link error"))
		return &pd.HeaderInfo{}, errors.New("fail")
	}

	exchange.AddHeader(&exchange.HeaderNodeInfo{
		Address:       request.Address,
		Port:          request.Port,
		NodeId:        request.NodeId,
		Weight:        request.Weight,
		CurrentTxId:   request.CurrentTxId,
		MasterAddress: request.MasterAddress,
		Service:       headerpd.NewDMQHeaderServiceClient(conn),
	})

	return &pd.HeaderInfo{}, nil
}
func (this *server) GetHeaderInfoRequest(ctx context.Context, request *pd.HeaderInfo) (out *pd.HeaderInfo, err error) {
	info := exchange.GetSelfHeaderInfo()
	out = &pd.HeaderInfo{}
	out.Address = info.Address
	out.Port = info.Port
	out.NodeId = info.NodeId
	out.Weight = info.Weight
	out.CurrentTxId = info.CurrentTxId
	out.MasterAddress = info.MasterAddress

	return out, nil
}
func (this *server) ProofClientRequest(ctx context.Context, request *pd.ProofClient) (out *pd.Response, err error) {
	key := request.Key
	out = &pd.Response{}
	str, err := mycrypto.Decrypt(filepath.Join(headerconfig.G_Crypto_Dir, "ppfile", "pri.pem"), key)
	if str != "gds" || err != nil {
		out.Errno = 1
		out.Errmsg = "密钥错误"
		return out, nil
	}
	file, ferr := os.Open(filepath.Join(headerconfig.G_Crypto_Dir, "keys"))
	if ferr != nil {
		out.Errno = 1
		out.Errmsg = "密钥错误"
		return out, nil
	}
	fileinfo, _ := file.Readdir(0)
	for _, v := range fileinfo {
		if v.Name() == key {
			out.Errno = 0
			out.Errmsg = "success"
			return out, nil
		}
	}
	out.Errno = 1
	out.Errmsg = "密钥错误"
	return out, nil
}

// follower 节点注册
func (this *server) FollowerRegistToHeaderRequest(ctx context.Context, request *pd.FollowerRegistToHeader) (out *pd.Response, err error) {
	fmt.Println(request, "返回流式数据")
	err1 := followerList.Add(request.GetAddress(), int(request.GetPort()))
	out = &pd.Response{}
	if err1 != nil {
		out.Errno = 1
		out.Errmsg = err1.Error()
	} else {
		out.Errno = 0
		out.Errmsg = "节点注册成功!"
		mylog.Info(fmt.Sprintf("%s:%d 节点注册成功", request.GetAddress(), int(request.GetPort())))
	}

	return
}

func writeDataToTopic(topic, message string) {
	if topicIsExist(topic) {
		topicfile, err := os.OpenFile(filepath.Join(headerconfig.G_Data_Dir, topic, "data"), os.O_EXCL|os.O_APPEND, 0655)
		if err != nil {
			return
		}
		topicMap[topic] = &Topic{file: topicfile, lock: new(sync.Mutex)}
	}

	t, ok := topicMap[topic]
	if !ok {
		return
	}
	t.lock.Lock()
	t.file.Seek(0, 2)
	t.file.Write(append([]byte(message), byte(0)))
	defer t.lock.Unlock()
}
func (this *server) FollowerYieldMsgDataRequest(in pd.DMQHeaderService_FollowerYieldMsgDataRequestServer) error {
	for info, _ := in.Recv(); info != nil; info, _ = in.Recv() {

		if !topicIsExist(info.Topic) {
			in.SendAndClose(&pd.Response{Errno: 1, Errmsg: "不存在的topic: " + info.Topic})
			return nil
		}
		info.Length = int64(len([]byte(info.Message)) + 1)
		EnterQueue(info)
		// 将数据写到dish
		// writeDataToTopic(info.Topic, info.Message)
		// info.Length = int64(len([]byte(info.Message)) + 1)
		// // 将数据添加到管道
		// msgChan <- *info
		// 管道
	}
	in.SendAndClose(&pd.Response{Errno: 2, Errmsg: "生产通道已关闭已关闭"})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// in.Send(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	// time.Sleep(time.Second * 4)
	return nil
}
func (this *server) FollowerToHeaderRequestDataRequest(in pd.DMQHeaderService_FollowerToHeaderRequestDataRequestServer) error {
	var address string
	var port int32
	flag := false
	for info, _ := in.Recv(); info != nil; info, _ = in.Recv() {
		if !flag {

			// 第一次是注册
			address = info.Address
			port = info.Port
			err := followerList.Add(address, int(port))
			if err != nil {
				in.Send(&pd.Response{Errno: 1, Errmsg: err.Error()})
				return nil
			}
			followerOutputPipe[fmt.Sprintf("%s_%d", address, port)] = &in
			flag = true
			mylog.Info(fmt.Sprintf("%s:%d 节点注册成功", address, int(port)))
			in.Send(&pd.Response{Errno: 0, Errmsg: "注册成功!"})
			continue
		}

		if !topicIsExist(info.Topic) {
			in.Send(&pd.Response{Errno: 1, Errmsg: "指定的topic不存在", Data: &pd.MessageData{Topic: info.Topic}})
			return nil
		}
		for i, _ := range followerList {
			if followerList[i].address == address && followerList[i].port == int(port) {
				if !followerList[i].HaveTopic(info.Topic) {
					followerList[i].AddTopic(info.Topic)
				}
			}
		}
		data, err := getMessageByGroupAndTopicAndOffset(info.Groupname, info.Topic, info.Offset)
		fmt.Println(data, err, info.Offset)
		if err == nil {
			for _, v := range data {
				fmt.Println("send: ", v, info.Groupname)
				in.Send(&pd.Response{Errno: 0, Errmsg: "success!", Data: &pd.MessageData{Des: info.Groupname, Topic: info.Topic, Message: v, Length: 1 + int64(len([]byte(v)))}})
			}

		} else {
			in.Send(&pd.Response{Errno: 1, Errmsg: err.Error(), Data: &pd.MessageData{Topic: info.Topic}})
		}
	}

	// 链接断开
	followerList.Remove(&FollowerNode{address: address, port: int(port)})
	delete(followerOutputPipe, fmt.Sprintf("%s_%d", address, port))
	mylog.Warning(fmt.Sprintf("%s:%d 节点已注销", address, int(port)))
	// in.SendAndClose(&pd.StudentResponseList{StudentResponse: []*pd.StudentResponse{&pd.StudentResponse{}}})
	return nil
}

func topicIsExist(topic string) bool {
	_, err := os.Open(filepath.Join(headerconfig.G_Data_Dir, topic))
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func topicGroupIsExist(topic, groupname string) bool {
	_, err := os.Open(filepath.Join(headerconfig.G_Data_Dir, topic, groupname))
	if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// 定时将 offset缓存写入到磁盘
// func writeCacheOffsetToDist() {
// 	time.Sleep(time.Second * time.Duration(headerconfig.G_Offset_Cache_Write_To_Dish_Time))
// 	var topic string
// 	var group string
// 	for p, v := range offsetCache {
// 		arr := strings.Split(p, "_")
// 		topic = arr[0]
// 		group = arr[1]
// 		file, err := os.OpenFile(filepath.Join(headerconfig.G_Data_Dir, topic, group, "offset"), os.O_TRUNC|os.O_EXCL, 0655)
// 		if err != nil {
// 			mylog.Error("offset缓存持久化错误: " + group + "->" + topic + " " + err.Error())
// 			continue
// 		}
// 		file.Write([]byte(fmt.Sprintf("%d", v)))
// 		file.Close()
// 	}
// }
func getConsumerGroupOffset(topicname, group string) int {
	file, err := os.Open(filepath.Join(headerconfig.G_Data_Dir, topicname, group, "offset"))
	defer file.Close()
	if err != nil {
		mylog.Error("position: header.go 145 " + err.Error())
		return 0
	}
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	offset, e := strconv.Atoi(string(line))
	if e != nil {
		offset = 0
	}
	return offset
}

// 当消费者第一次注册时获取数据
func getMessageByGroupAndTopicAndOffset(group string, topicname string, groupOffset int64) ([]string, error) {
	if !topicIsExist(topicname) {
		return []string{}, errors.New("指定的 topic 不存在!")
	}
	/*
		{
			topicfile
			lock
		}
	*/
	topic, ok := topicMap[topicname]

	// 当该topic第一次被消费的时候
	if !ok {
		topicfile, err := os.OpenFile(filepath.Join(headerconfig.G_Data_Dir, topicname, "data"), os.O_EXCL|os.O_APPEND, 0655)
		if err != nil {

			// 话题下的数据文件不存在
			mylog.Info(topicname + " 数据读取失败: " + err.Error())
			return []string{}, errors.New("话题数据读取异常!")
		}
		topicMap[topicname] = &Topic{file: topicfile, lock: new(sync.Mutex)}
		// 将偏移量缓存在内存
		// offsetCache[topicname+"_"+group] = offset

		topicMap[topicname].lock.Lock()
		data, _ := getDataByFileAndOffset(topicfile, groupOffset)
		topicMap[topicname].lock.Unlock()

		// 将新的偏移量记录在缓存中
		// offsetCache[topicname+"_"+group] = newoffset
		return data, nil
	} else {
		topicfile := topic.file
		// 将偏移量缓存在内存
		// offsetCache[topicname+"_"+group] = offset

		topic.lock.Lock()
		data, _ := getDataByFileAndOffset(topicfile, groupOffset)
		topic.lock.Unlock()

		// 将新的偏移量记录在缓存中
		// offsetCache[topicname+"_"+group] = newoffset
		return data, nil
	}
	return []string{}, nil
}

// 缓存数据，以及当前偏移量
func getDataByFileAndOffset(file *os.File, offset int64) ([]string, int) {
	se, _err := file.Seek(int64(offset), 0)
	if _err != nil {
		return []string{}, int(se)
	}
	msglist := make([]string, 0)
	cache := make([]byte, 1024)
	pre := make([]byte, 0)
	var err error
	var n int
	n = 0
	for {
		n, err = file.Read(cache)
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			// fmt.Println(cache[i])
			if cache[i] == byte(0) {
				// 发现分隔符
				msglist = append(msglist, string(pre))

				pre = make([]byte, 0)
				continue
			}
			pre = append(pre, cache[i])
		}
	}
	if len(pre) != 0 {
		msglist = append(msglist, string(pre))
	}
	newoffset, _ := file.Seek(0, 1)
	return msglist, int(newoffset)
}

func channelAroused() {
	var msgdata pd.MessageData
	for {
		msgdata = <-msgChan
		for {
			if len(followerList) <= 0 {
				break
			}

			start := rand.Intn(len(followerList))
			index := start
			for {
				if len(followerList) <= 0 {
					break
				}
				node := followerList[index]
				if node.HaveTopic(msgdata.Topic) {
					in, ok := followerOutputPipe[fmt.Sprintf("%s_%d", node.address, node.port)]
					if !ok {
						continue
					}
					(*in).Send(&pd.Response{Errno: 0, Errmsg: "success!", Data: &msgdata})
					break
				}
				index++
				index = index % len(followerList)
				if index == start {
					break
				}
			}
			break

		}
	}
}
func initGlobalVar() {
	// 初始化 followerList 列表(初始为空)
	followerList = make(FollowerNodeList, 0)
	topicMap = make(map[string]*Topic, 0)

	// map[topic+group] = offset
	// offsetCache = make(map[string]int, 0)
	topicLock = make(map[string]*sync.Mutex, 0)

	followerOutputPipe = make(map[string]*pd.DMQHeaderService_FollowerToHeaderRequestDataRequestServer, 0)
	msgChan = make(chan pd.MessageData, 100)
}

func StartHeader(conf *utils.MyConfig) {
	initGlobalVar()
	tx.Init(conf)
	exchange.Init(conf)

	headerconfig = conf
	mylog.Info(fmt.Sprintln("正在启动 header 节点...."))
	//创建网络
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.G_Bind_Address, conf.G_Bind_Port))
	if err != nil {
		fmt.Println("网络错误", err)
	}

	//创建grpc的服务
	srv := grpc.NewServer()

	//注册服务
	pd.RegisterDMQHeaderServiceServer(srv, &server{})
	mylog.Info(fmt.Sprintf("listen: %s:%d\n", conf.G_Bind_Address, conf.G_Bind_Port))

	go startWork()
	//等待网络连接
	err = srv.Serve(ln)
	if err != nil {
		mylog.Error("启动失败: " + err.Error())
	}

}
func startWork() {
	time.Sleep(time.Second * 3)
	exchange.StartExchange()
	go channelAroused()
}
