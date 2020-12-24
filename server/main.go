package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"server/mycmd"

	"server/mylog"
	"server/service"
	"server/utils"
)

var config *utils.MyConfig

func main() {
	// 解析cli
	cmd := mycmd.NewCmd()
	configfile, err := cmd.Get("configfile")
	cmdstr, _ := cmd.Get("useCmd")
	if err != nil {
		mylog.Error("没有指定配置文件!")
		os.Exit(0)
	}

	// 记载配置文件
	config = utils.InitConfig(configfile)
	if cmdstr == "true" {
		startCmdLine()
		return
	}

	if config.G_Node_Type == "header" {
		mylog.Info("正在启动 header 节点......")
		service.StartHeaderService(config)
	} else if config.G_Node_Type == "follower" {
		mylog.Info("正在启动 follower 节点......")
		service.StartFollowerService(config)
	} else {
		mylog.Error("config: 错误的节点类型!")
	}
}
func startCmdLine() {
	var cmd string
	for {
		fmt.Print("DMQ> ")
		reader := bufio.NewReader(os.Stdin)
		strBytes, _, _ := reader.ReadLine()
		cmd = string(strBytes)
		if strings.TrimSpace(cmd) == "" {
			continue
		}
		var arg1 string
		var arg2 string
		arr := strings.Split(cmd, " ")
		arg1 = arr[0]
		if len(arr) > 3 {
			mylog.Error("错误的指示")
			continue
		}
		if len(arr) == 2 {
			arg2 = arr[1]
		}
		switch arg1 {
		case "help":
			fmt.Println("create			创建一个新的topic")
			fmt.Println("delete			创建一个新的topic")
			fmt.Println("list			查看 topic 列表")
			fmt.Println("clean			清空指定 topic 数据")
			fmt.Println("exit			退出终端")
		case "create":
			err := os.Mkdir(filepath.Join(config.G_Data_Dir, arg2), 0755)
			if err != nil {
				mylog.Error("topic创建失败: 已存在的topic")
				break
			}
			file, _err := os.Create(filepath.Join(config.G_Data_Dir, arg2, "data"))
			file.Close()
			if _err != nil {
				mylog.Error("创建失败: " + err.Error())
				break
			}
			mylog.Success(fmt.Sprintf("创建成功: %s", arg2))
		case "delete":
			err := os.RemoveAll(filepath.Join(config.G_Data_Dir, arg2))
			if err != nil {
				mylog.Error(fmt.Sprintf("topic(%s) 删除失败: %s", arg2, err.Error()))
				break
			}
			mylog.Success(fmt.Sprintf("topic(%s) 删除成功", arg2))
		case "list":
			files, err := ioutil.ReadDir(config.G_Data_Dir)
			if err != nil {
				mylog.Error("error: " + err.Error())
			}
			for _, v := range files {
				fmt.Println(v.Name())
			}
		case "exit":
			os.Exit(0)
		}
	}
}
