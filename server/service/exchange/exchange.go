package exchange

import (
	"context"
	"fmt"
	"os"
	"server/mylog"
	headerpd "server/service/proto/header"
	"server/service/tx"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type HeaderNodeInfo struct {
	Address       string
	Port          int32
	NodeId        int32
	Weight        int32
	CurrentTxId   string
	Service       headerpd.DMQHeaderServiceClient
	MasterAddress string
}

// 集群中的header节点
var headers []*HeaderNodeInfo = make([]*HeaderNodeInfo, 0, 0)

var selfHeaderInfo *HeaderNodeInfo = &HeaderNodeInfo{}

/*
cluster_status
	vote  正在选举
	operation  正在运行
*/
var cluster_status string
var current_master *HeaderNodeInfo

var config *utils.MyConfig

func GetCurrentMaster() *HeaderNodeInfo {
	return current_master
}
func GetCluterStatus() string {
	return cluster_status
}
func GetHeaders() []*HeaderNodeInfo {
	return headers
}
func AddHeader(newheader *HeaderNodeInfo) {
	headers = append(headers, newheader)
}

// 开始选举
func StartVode() {
	cluster_status = "vote"

}
func GetSelfHeaderInfo() *HeaderNodeInfo {
	return selfHeaderInfo
}
func GetQuorumInfo() []*HeaderNodeInfo {
	headers = make([]*HeaderNodeInfo, 0)
	_headers := make([]*HeaderNodeInfo, 0)
	for _, v := range config.G_Quorum {
		addr := strings.Split(v, ":")
		if len(addr) != 2 {
			mylog.Error("quroums error")
			os.Exit(0)
		}
		port, _ := strconv.Atoi(addr[1])
		_headers = append(_headers, &HeaderNodeInfo{Address: addr[0], Port: int32(port)})
	}
	for i := 0; i < len(_headers); i++ {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", _headers[i].Address, _headers[i].Port), grpc.WithInsecure())
		if err != nil {
			mylog.Error(fmt.Sprintf("%s:%d  %s", _headers[i].Address, _headers[i].Port, "link error"))
			continue
		}
		_headers[i].Service = headerpd.NewDMQHeaderServiceClient(conn)
		info, err1 := _headers[i].Service.GetHeaderInfoRequest(context.Background(), &headerpd.HeaderInfo{
			Address:     selfHeaderInfo.Address,
			Port:        int32(selfHeaderInfo.Port),
			NodeId:      int32(selfHeaderInfo.NodeId),
			Weight:      int32(selfHeaderInfo.Weight),
			CurrentTxId: selfHeaderInfo.CurrentTxId,
		})
		if err1 != nil {
			mylog.Error(fmt.Sprintf("%s:%d  %s", _headers[i].Address, _headers[i].Port, "link error"))
			continue
		}
		_headers[i].NodeId = info.NodeId
		_headers[i].Weight = info.Weight
		_headers[i].CurrentTxId = info.CurrentTxId
		_headers[i].MasterAddress = info.MasterAddress

		if _headers[i].MasterAddress != "" {
			selfHeaderInfo.MasterAddress = _headers[i].MasterAddress
		}
		headers = append(headers, _headers[i])

	}
	return headers
}
func GetMaster() {
	if selfHeaderInfo.MasterAddress != "" {
		// 集群中已存在master
		for i := 0; i < len(headers); i++ {
			if fmt.Sprintf("%s:%d", headers[i].Address, headers[i].Port) == selfHeaderInfo.MasterAddress {
				current_master = headers[i]
				break
			}
		}
		if selfHeaderInfo.MasterAddress == fmt.Sprintf("%s:%d", selfHeaderInfo.Address, selfHeaderInfo.Port) {
			current_master = selfHeaderInfo
		}
	} else {
		// 选举master

		// 更新集群状态,改状态将禁止所有的请求(以后可以优化这里,"优雅的选举")
		cluster_status = "vote"

		// 重新获取各节点信息排除故障节点
		GetHeaders()
		if len(headers) == 0 {
			current_master = selfHeaderInfo
			selfHeaderInfo.MasterAddress = fmt.Sprintf("%s:%d", current_master.Address, current_master.Port)
			cluster_status = "operation"
			return
		}
		candidates1 := make([]*HeaderNodeInfo, 0, 0)
		for i := 0; i < len(headers); i++ {
			if len(candidates1) == 0 {
				candidates1 = append(candidates1, headers[i])
			} else if tx.CompareTxId(headers[i].CurrentTxId, candidates1[0].CurrentTxId) == 1 {
				candidates1 = make([]*HeaderNodeInfo, 0, 0)
			}
			candidates1 = append(candidates1, headers[i])
		}

		// 胜出(该节点事务id最新)
		candidates2 := make([]*HeaderNodeInfo, 0, 0)
		if len(candidates1) == 1 {
			current_master = candidates1[0]
			selfHeaderInfo.MasterAddress = fmt.Sprintf("%s:%d", current_master.Address, current_master.Port)
			cluster_status = "operation"
			return
		}

		for i := 0; i < len(candidates1); i++ {
			if len(candidates2) == 0 {
				candidates2 = append(candidates2, candidates1[i])
			} else if candidates1[i].Weight > candidates2[0].Weight {
				candidates2 = make([]*HeaderNodeInfo, 0, 0)
			}
			candidates2 = append(candidates2, candidates1[i])
		}

		// 胜出(该节点权重最高)
		if len(candidates2) == 1 {
			current_master = candidates2[0]
			selfHeaderInfo.MasterAddress = fmt.Sprintf("%s:%d", current_master.Address, current_master.Port)
			cluster_status = "operation"

			return
		}
		current_master = candidates2[0]
		for i := 1; i < len(candidates2); i++ {
			if candidates2[i].NodeId > current_master.NodeId {
				current_master = candidates2[i]
			}
		}
		selfHeaderInfo.MasterAddress = fmt.Sprintf("%s:%d", current_master.Address, current_master.Port)
		cluster_status = "operation"

	}
}
func broadcastHeaderEnter() {
	GetQuorumInfo()
	for i := 0; i < len(headers); i++ {
		headers[i].Service.EnterCluster(context.Background(), &headerpd.HeaderInfo{
			Address:     selfHeaderInfo.Address,
			Port:        int32(selfHeaderInfo.Port),
			NodeId:      int32(selfHeaderInfo.NodeId),
			Weight:      int32(selfHeaderInfo.Weight),
			CurrentTxId: selfHeaderInfo.CurrentTxId,
		})
	}
}
func Init(conf *utils.MyConfig) {
	config = conf

	selfHeaderInfo.NodeId = int32(config.G_Node_Id)
	selfHeaderInfo.Weight = int32(config.G_Weight)
	selfHeaderInfo.CurrentTxId = tx.GetCurrentTxId()
	selfHeaderInfo.Address = config.G_Bind_Address
	selfHeaderInfo.Port = int32(config.G_Bind_Port)
	selfHeaderInfo.MasterAddress = ""

}
func StartExchange() {
	GetQuorumInfo()
	// 获取master
	GetMaster()
	fmt.Println(current_master, "===================")
	broadcastHeaderEnter()
	mylog.Info(fmt.Sprintf("当前master: %s:%d", current_master.Address, current_master.Port))
	fmt.Println("集群成员")
	go p()

}
func p() {
	for _, v := range headers {
		fmt.Println(v.Address, v.Port)
	}
	time.Sleep(time.Second)
	p()
}