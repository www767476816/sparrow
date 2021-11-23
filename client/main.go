package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//func  testrpc()  {
//	// 1. 新建连接，端口是服务端开放的8082端口
//	// 并且添加grpc.WithInsecure()，不然没有证书会报错
//	client :=rpc_service.CreateClient("127.0.0.1","2013")
//	client.Connect()
//
//	// 退出时关闭链接
//	defer client.Close()
//
//	// 2. 调用Product.pb.go中的NewProdServiceClient方法
//	productServiceClient := service.NewProdServiceClient(client.GetConnect())
//
//	// 3. 直接像调用本地方法一样调用GetProductStock方法
//	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 233})
//	if err != nil {
//		log.Fatal("调用gRPC方法错误: ", err)
//	}
//
//	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
//}
func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1",
		url.Values{"name": {"www"}, "old": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
func main() {

}
