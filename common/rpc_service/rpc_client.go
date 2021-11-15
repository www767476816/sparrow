package rpc_service

import (
	"google.golang.org/grpc"
)

type RpcClient struct {
	clientConnect *grpc.ClientConn
	serverAddr string
	serverPort string
}
//创建客户端对象
func CreateClient(addr string,port string) *RpcClient{
	return &RpcClient{nil,addr,port}
}
//连接
func (c *RpcClient)Connect() error{
	var err error
	c.clientConnect, err = grpc.Dial(c.serverAddr+":"+c.serverPort, grpc.WithInsecure())

	if err != nil {
		c.Close()
	}
	return err
}
//关闭
func (c *RpcClient)Close(){
	c.clientConnect.Close()
	c.clientConnect=nil
	c.serverAddr=""
	c.serverPort=""
}
//获取rpc客户端
func (c *RpcClient)GetConnect() *grpc.ClientConn{
	return c.clientConnect
}
//获取服务器地址
func (c *RpcClient)GetAddr() string{
	return c.serverAddr
}
//获取服务器端口
func (c *RpcClient)GetPort() string{
	return c.serverPort
}