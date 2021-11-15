package tcp_service

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)
const MAX_DATA_LEN=65535
type NetInfo struct {
	port string
	listener net.Listener
	connectCount int32
	maxConnect int32
}

func NewTCPService(port string,maxConnect int32) *NetInfo {
	return &NetInfo{
		port :port,
		connectCount :0,
		maxConnect:maxConnect,
	}
}

func (s *NetInfo) Init() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	s.listener=listener
	return nil
}
func (s *NetInfo) Start()error{
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.receive(conn)
	}
	return nil
}
func (s *NetInfo) receive(conn net.Conn)error{
	defer conn.Close()
	for{
		sizeByte :=make([]byte,4)
		if _,err :=conn.Read(sizeByte);err != nil && err != io.EOF{
			return err
		}
		size :=binary.BigEndian.Uint32(sizeByte)
		if size>MAX_DATA_LEN-4{
			return errors.New("receive data len is error,max:"+string(MAX_DATA_LEN)+",recv:"+string(size))
		}
		dataByte :=make([]byte,size)
		if _,err :=conn.Read(dataByte);err != nil && err != io.EOF{
			return err
		}
		data :=string(dataByte)
		fmt.Println("接收到数据：",data)
	}
}

func (s *NetInfo) Send(conn net.Conn,data string) error{
	if len(data)>MAX_DATA_LEN-4{
		return errors.New("Send data len is error,max:"+string(MAX_DATA_LEN)+",recv:"+string(len(data)))
	}
	head :=make([]byte,4)
	head[3]=uint8(len(data))
	head[2]=uint8(len(data)>>8)
	head[1]=uint8(len(data)>>16)
	head[0]=uint8(len(data)>>24)

	go conn.Write(append(head,[]byte(data)...))
	return nil
}
