package main

import (
	"fmt"
	"net/http"
	"time"

	//草拟吗 自己创建的目录 哈哈哈哈哈    还好我比较聪明  要不然 就完蛋了  麻痹
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/shiming", websocket.Handler(Echo))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func Echo(w *websocket.Conn) {
	fmt.Println("...............................")
	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		msg := "我已经收到消息 Received:" + reply
		//  连接的话 只能是   string；类型的啊
		fmt.Println("发给客户端的消息： " + msg)

		go func() {
			time.Sleep(time.Second * 4)
			if error = websocket.Message.Send(w, msg); error != nil {
				fmt.Println("不能够发送消息 悲催哦")

			}
		}()
	}
}
