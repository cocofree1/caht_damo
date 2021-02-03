package client

import (
	"github.com/opentracing/opentracing-go/log"
	"net"
)

type Connection interface {
	GetConnect()net.Conn
}

type NetConnection struct {
	netConn  net.Conn
}

func NewNetConnection(con Connection)*NetConnection{
	return &NetConnection{netConn: con.GetConnect()}
}

func (n* NetConnection)SendMessage(msg []byte){
	n.netConn.Write(msg)
}

func (n* NetConnection)ReadMessage()string{
	var buff [128]byte
	byteLen,err := n.netConn.Read(buff[:])
	if err != nil{
		log.Error(err)
	}
	return string(buff[:byteLen])
}

func (n* NetConnection)Close(){
	n.netConn.Close()
}