package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"server/mylog"
	"server/utils/config"
)

type MyConfig struct {
	G_Node_Type                       string // 节点类型
	G_Header_Address                  string
	G_Current_Platform                string // 当前 os
	G_Header_Port                     int
	G_Data_Dir                        string // 数据存储目录(header节点为必填)
	G_Crypto_Dir                      string // auth路径
	G_Heartbeat                       int    // 心跳检测间隔 header -> follower
	G_Bind_Address                    string
	G_Bind_Port                       int
	G_Cache_Address                   string
	G_Cache_Port                      int
	G_Max_Register                    int // follower 最大注册消费者数
	G_Ws_Address                      string
	G_Node_Id                         int
	G_Http_Address                    string
	G_Http_Port                       int
	G_Weight                          int
	G_Quorum                          []string
	G_Tx_Dir                          string
	G_Offset_Cache_Write_To_Dish_Time int
}

func InitConfig(configpath string) *MyConfig {
	myconfig := MyConfig{}
	var err error

	//从配置文件读取配置信息
	//如果项目迁移需要进行修改
	appconf, err := config.NewConfig("ini", configpath)
	if err != nil {
		mylog.Error(err.Error())
		os.Exit(0)
	}
	myconfig.G_Node_Type = appconf.String("node_type")
	// mylog.Success(G_Node_Type)
	if myconfig.G_Node_Type != "header" && myconfig.G_Node_Type != "follower" {
		// 必须明确节点类型 header and follower
		mylog.Error("config: node_type option not allow empty! ought input value that items 'header' or 'follower'")
		os.Exit(0)
	}
	myconfig.G_Current_Platform = runtime.GOOS
	myconfig.G_Ws_Address = appconf.String("ws_addr") + ":" + appconf.String("ws_port")
	if myconfig.G_Node_Type == "follower"{
		myconfig.G_Header_Address = appconf.String("header_address")
		myconfig.G_Header_Port, err = appconf.Int("header_port")
		if err != nil {
			mylog.Error("config: " + err.Error())
			os.Exit(0)
		}
	}

	if myconfig.G_Node_Type == "header" {
		myconfig.G_Data_Dir = appconf.String("data_dir")

		_, err = os.Open(myconfig.G_Data_Dir)
		if os.IsNotExist(err) {
			mylog.Error("config: " + myconfig.G_Data_Dir + " not found!")
			os.Exit(0)
		}

		myconfig.G_Crypto_Dir = appconf.String("crypto_dir")

		_, err = os.Open(myconfig.G_Crypto_Dir)
		if os.IsNotExist(err) {
			mylog.Error("config: " + myconfig.G_Crypto_Dir + " not found!")
			os.Exit(0)
		}

		myconfig.G_Tx_Dir = appconf.String("tx_dir")

		_, err = os.Open(myconfig.G_Tx_Dir)
		if os.IsNotExist(err) {
			os.Mkdir(myconfig.G_Tx_Dir, 0755)
			file, _ := os.Create(filepath.Join(myconfig.G_Tx_Dir, "finish"))
			file.Close()
			file, _ = os.Create(filepath.Join(myconfig.G_Tx_Dir, "unfinish"))
			file.Close()
		}

		myconfig.G_Node_Id, err = appconf.Int("node_id")
		if err != nil {
			mylog.Error("config: " + err.Error())
			os.Exit(0)
		}
		myconfig.G_Weight, err = appconf.Int("weight")
		if err != nil {
			mylog.Error("config: " + err.Error())
			os.Exit(0)
		}

		str := appconf.String("quorum")
		if strings.TrimSpace(str) == "" {
			mylog.Error("config: " + "quorum is empty!")
			os.Exit(0)
		}
		myconfig.G_Quorum = strings.Split(str, ",")

	}
	if myconfig.G_Node_Type == "follower" {
		myconfig.G_Cache_Address = appconf.String("cache_address")
		myconfig.G_Cache_Port, err = appconf.Int("cache_port")
		if err != nil {
			mylog.Error("config: " + err.Error())
		}

		myconfig.G_Http_Address = appconf.String("http_addr")
		myconfig.G_Http_Port, err = appconf.Int("http_port")
		if err != nil {
			mylog.Error("config: " + err.Error())
		}
	}

	myconfig.G_Heartbeat, err = appconf.Int("heartbeat")
	if err != nil {
		mylog.Error("config: " + err.Error())
	}
	myconfig.G_Bind_Address = appconf.String("bind_address")
	myconfig.G_Bind_Port, err = appconf.Int("bind_port")
	if err != nil {
		mylog.Error("config: " + err.Error())
	}
	myconfig.G_Max_Register, err = appconf.Int("max_register")
	if err != nil {
		mylog.Error("config: " + err.Error())
	}

	myconfig.G_Offset_Cache_Write_To_Dish_Time, err = appconf.Int("offset_cache_write_to_dish_time")
	if err != nil {
		mylog.Error("config: " + err.Error())
	}
	return &myconfig
}
