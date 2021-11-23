package main

import (
	"encoding/xml"
	"fmt"
	_ "github.com/go_protocol-sql-driver/mysql"
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

//func (ps *ProdService) GetProductStock(ctx context.Context, request *rpc_protocol.ProductRequest) (*rpc_protocol.ProductResponse, error) {
//	return &rpc_protocol.ProductResponse{ProdStock: request.ProdId*2}, nil
//}

//func testRpc(){
//	server :=rpc_service.CreateServer("tcp","2013")
//	fmt.Println(server.GetNet(),";",server.GetPort())
//	server.Init()
//
//
//	server.Start()
//	rpc_protocol.RegisterProdServiceServer(server.GetConnect(), new(ProdService))
//	server.Run()
//
//}

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
	testXml()
}
