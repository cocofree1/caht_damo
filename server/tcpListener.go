package server

import (
	"github.com/opentracing/opentracing-go/log"
	"net"
)

type TcpListener struct{
	hostAddr    string
}

func NewTcpListener(HostAddr string)Listener{
	return &TcpListener{hostAddr: HostAddr}
}

func (t* TcpListener)GetListener()net.Listener{
	listener,err := net.Listen("tcp",t.hostAddr)
	if err != nil{
		log.Error(err)
	}
	return listener
}
