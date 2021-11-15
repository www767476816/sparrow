package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
	"io/ioutil"
	"os"
	"os/signal"
	"rpc-demo/common/rpc_service"
	"rpc-demo/common/rpc_service/service"
	"strconv"
	//"rpc-demo/common/sql_service"
	//"time"
)

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *service.ProductRequest) (*service.ProductResponse, error) {
	return &service.ProductResponse{ProdStock: request.ProdId*2}, nil
}

func testRpc(){
	server :=rpc_service.CreateServer("tcp","2013")
	fmt.Println(server.GetNet(),";",server.GetPort())
	server.Init()


	server.Start()
	service.RegisterProdServiceServer(server.GetConnect(), new(ProdService))
	server.Run()

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
	var ps BookStore
	if erro :=xml.Unmarshal(result,&ps);erro!=nil{
		fmt.Println(erro)
		return
	}
	fmt.Println(ps)
	//fmt.Println(ps.root,ps.min,ps.max)
}

func main() {
	err :=errors.New("invalid service type:"+strconv.Itoa(1))
	fmt.Println(err)
}
