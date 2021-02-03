package chat

import (
	"chat_damo/server"
	"encoding/json"
	"fmt"
	"net"
)

type Chat struct{
	name   string
	conn   net.Conn
}

var chatMap = make(map[string]*Chat)

func ServerSingleChat(){
	netListen := ServerChat()
	defer func() {
		netListen.Close()
	}()
	for {
		// 每连接注册一个
		conn := netListen.NetAccept()
		name := fmt.Sprintf("kkk%s",conn.RemoteAddr().String())
		chat := &Chat{
			name: name,
			conn: conn,
		}
		chatMap[name] = chat
		// 读取消息
		fmt.Println(name," 已上线")
		go DealServerMessage(netListen,conn)
	}
}

func ServerChat()*server.NetListener{
	tcpListen := server.NewTcpListener("127.0.0.1:3400")
	netListen := server.NewNetListener(tcpListen)
	return netListen
}

func DealServerMessage(netListen *server.NetListener,conn net.Conn){
	for {
		readMsg := netListen.ReadMessage(conn)
		if len(readMsg) > 0 {
			var msg Message
			json.Unmarshal(readMsg,&msg)
			// 群聊就是遍历map向所有人发送消息
			netListen.SendMessage(msg.Msg,chatMap[msg.Name].conn)
		}
	}
}