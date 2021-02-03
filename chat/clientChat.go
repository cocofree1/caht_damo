package chat

import (
	"chat_damo/client"
	"chat_damo/lib"
	"encoding/json"
	"fmt"
	"strings"
)

type Message struct{
	Name  string
	Msg   string
}

func ClientSingleChat(netClient *client.NetConnection){
	go SendSingleMessage(netClient)
	go ReadMessage(netClient)
}

func ClientChat()*client.NetConnection{
	tcpConn := client.NewTcpConnection("127.0.0.1:3400")
	netClient := client.NewNetConnection(tcpConn)
	return netClient
}

func SendSingleMessage(netClient *client.NetConnection){
	cmdReader := lib.NewCmdReader()
	for {
		msg := cmdReader.ReadCmd()
		if len(msg) > 0{
			msgArr := strings.Split(msg,"#")
			chatMsg := Message{
				Name: msgArr[0],
				Msg: msgArr[1],
			}
			chatByte,_ := json.Marshal(chatMsg)
			netClient.SendMessage(chatByte)
		}
	}
}

func ReadMessage(netClient *client.NetConnection){
	for {
		msg := netClient.ReadMessage()
		if len(msg) > 0 {
			fmt.Println(msg)
		}
	}
}