package main

import "chat_damo/chat"

func main(){
	quit := make(chan int, 0)
	chat.ServerSingleChat()
	<- quit
}
