package client

import (
"github.com/opentracing/opentracing-go/log"
"net"
)

type TcpConnection struct{
	hostAddr    string
}

func NewTcpConnection(HostAddr string)Connection{
	return &TcpConnection{hostAddr: HostAddr}
}

func (t* TcpConnection)GetConnect()net.Conn{
	con,err := net.Dial("tcp",t.hostAddr)
	if err != nil{
		log.Error(err)
	}
	return con
}
