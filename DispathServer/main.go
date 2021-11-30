package main

import (
	"encoding/xml"
	"fmt"
	"rpc-demo/DispathServer/rpc_package"

	//_ "github.com/go_protocol-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc-demo/common/frame"
	//google_protobuf "google.golang.org/protobuf/types/known/emptypb"
	"io/ioutil"
	"os"
	"os/signal"

	//"rpc-demo/common/sql_service"
	//"time"
)

type ProdService struct {
}

func (ps *ProdService) SayHello(ctx context.Context, request *rpc_package.HelloRequest) (*rpc_package.HelloReply, error) {
	return &rpc_package.HelloReply{Message: "nihao"}, nil
}

func testRpc(){
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()

	// 2. 将刚刚我们新建的ProdService注册进去
	rpc_package.RegisterHelloWorldServiceServer(rpcServer, new(ProdService))

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	runErr := rpcServer.Serve(listener)
	fmt.Println(runErr)

}

func close(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("quit,Got signal:", s)
}

type BookStore struct {
		 XMLName  xml.Name `xml:"bookstore"`
	     Name int32 `xml:"bookstore"`
}
func testXml(){
	//testMysql()
	//close()
	result,err :=ioutil.ReadFile("./DispathServer/config.xml")
	if err!=nil{
		fmt.Println(err)
		return
	}
	var ps frame.NormalConfig
	if erro :=xml.Unmarshal(result,&ps);erro!=nil{
		fmt.Println(erro)
		return
	}
	fmt.Println(ps)
	//fmt.Println(ps.root,ps.min,ps.max)
}

func main() {
	testRpc()
}
