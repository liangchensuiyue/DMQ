package mycmd

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/mycrypto"
	"server/mylog"
	"server/utils"
	"strconv"
	"strings"
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

func (mycmd *Mycmd) StartCmdLine(config *utils.MyConfig) {
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
			fmt.Println("listtopic			查看 topic 列表")
			fmt.Println("clean			清空指定 topic 数据")
			fmt.Println("listkeys		   查看密钥列表")
			fmt.Println("deletekey			根据索引删除密钥")
			fmt.Println("newkey			产生新的密钥")
			fmt.Println("resetcrypto		重新生产密钥系统(之前的密钥都失效)")
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
		case "listtopic":
			files, err := ioutil.ReadDir(config.G_Data_Dir)
			if err != nil {
				mylog.Error("error: " + err.Error())
			}
			for _, v := range files {
				fmt.Println(v.Name())
			}
		case "resetcrypto":
			os.RemoveAll(filepath.Join(config.G_Crypto_Dir, "keys"))
			os.Mkdir(filepath.Join(config.G_Crypto_Dir, "ppfile"), 0755)
			os.Mkdir(filepath.Join(config.G_Crypto_Dir, "keys"), 0755)
			mycrypto.GenerateRsaKey(256, filepath.Join(config.G_Crypto_Dir, "ppfile"))
			mylog.Success("success")
		case "listkeys":
			file, err := os.Open(filepath.Join(config.G_Crypto_Dir, "keys"))
			if err != nil {
				mylog.Error("error: " + err.Error())
				break
			}
			fileinfo, _ := file.Readdir(0)
			for k, v := range fileinfo {
				fmt.Println(k, v.Name())
			}
		case "newkey":
			newkey, err := mycrypto.Encrypt(filepath.Join(config.G_Crypto_Dir, "ppfile", "pub.pem"), "gds")
			if err != nil {
				mylog.Error("error: " + err.Error())
				break
			}
			fmt.Println(newkey)
			file, _ := os.Create(filepath.Join(config.G_Crypto_Dir, "keys", newkey))
			file.Close()
			mylog.Success("success")
		case "deletekey":
			index, err := strconv.Atoi(arg2)
			if err != nil {
				mylog.Error("error: " + err.Error())
				break
			}
			file, err := os.Open(filepath.Join(config.G_Crypto_Dir, "keys"))
			if err != nil {
				mylog.Error("error: " + err.Error())
				break
			}
			fileinfo, _ := file.Readdir(0)
			if index >= len(fileinfo) {
				break
			}
			err = os.Remove(filepath.Join(config.G_Crypto_Dir, "keys", fileinfo[index].Name()))
			if err != nil {
				mylog.Error("error: " + err.Error())
				break
			} else {
				mylog.Success("success")
				break
			}

		case "exit":
			os.Exit(0)
		}
	}
}
