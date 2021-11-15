package rpc_service

import (
	"google.golang.org/grpc"
	"net"
)

type RpcServer struct {
	serverConn *grpc.Server
	listener net.Listener
	netType string
	port string
}
//创建服务器对象
func CreateServer(netType string,port string) *RpcServer{

	return &RpcServer{nil,nil,netType,port}
}
//初始化服务器
func (s *RpcServer)Init(){
	s.serverConn=grpc.NewServer()
}
//开始监听端口
func (s *RpcServer)Start() error {
	var err error
	s.listener,err=net.Listen(s.netType, ":"+s.port)
	if err !=nil{
		s.Close()
	}
	return err
}
//运行
func (s *RpcServer)Run() error{
	return s.serverConn.Serve(s.listener)
}
//关闭
func (s *RpcServer)Close() {
	s.serverConn.Stop()
	s.serverConn=nil
	s.listener=nil
	s.netType=""
	s.port=""
}
//获取rpc连接
func (s *RpcServer)GetConnect() *grpc.Server{
	return s.serverConn
}
//获取listener
func (s *RpcServer)GetListener() net.Listener{
	return s.listener
}
//获取net类型
func (s *RpcServer)GetNet() string{
	return s.netType
}
//获取端口
func (s *RpcServer)GetPort() string{
	return s.port
}
