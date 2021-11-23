package rpc_service

import (
	"google.golang.org/grpc"
)

type RpcClient struct {
	clientConnect *grpc.ClientConn
	serverType uint32
	serverIP string
	serverPort string
}
//创建客户端对象
func CreateClient(addr string,port string,serverType uint32) *RpcClient{
	return &RpcClient{nil,serverType,addr,port}
}
//连接
func (c *RpcClient)Connect() error{
	var err error
	c.clientConnect, err = grpc.Dial(c.serverIP+":"+c.serverPort, grpc.WithInsecure())

	if err != nil {
		c.Close()
	}
	return err
}
//关闭
func (c *RpcClient)Close(){
	c.clientConnect.Close()
	c.clientConnect=nil
	c.serverIP=""
	c.serverPort=""
}
//获取rpc客户端
func (c *RpcClient)GetConnect() *grpc.ClientConn{
	return c.clientConnect
}
//获取服务器地址
func (c *RpcClient)GetIP() string{
	return c.serverIP
}
//获取服务器类型
func (c *RpcClient)GetPort() string{
	return c.serverPort
}
func (c *RpcClient)GetServerType() uint32{
	return c.serverType
}