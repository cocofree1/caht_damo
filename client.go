package main

import (
	"chat_damo/chat"
)

func main(){
	quit := make(chan int, 0)
	netClient := chat.ClientChat()
	defer func() {
		netClient.Close()
	}()
	chat.ClientSingleChat(netClient)
	<- quit
}