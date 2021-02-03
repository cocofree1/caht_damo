package server

import (
	"github.com/opentracing/opentracing-go/log"
	"net"
)

type Listener interface {
	GetListener()net.Listener
}

type NetListener struct {
	listen  net.Listener
}

func NewNetListener(listener Listener)*NetListener{
	return &NetListener{listen: listener.GetListener()}
}

func (n* NetListener)NetAccept()net.Conn{
	con,err := n.listen.Accept()
	if err != nil{
		log.Error(err)
	}
	return con
}

func (n* NetListener)SendMessage(msg string, netConn net.Conn){
	netConn.Write([]byte(msg))
}

func (n* NetListener)ReadMessage(netConn net.Conn)[]byte{
	var buff [128]byte
	msgLen,err := netConn.Read(buff[:])
	if err != nil{
		log.Error(err)
	}
	return buff[:msgLen]
}

func (n* NetListener)Close(){
	n.listen.Close()
}