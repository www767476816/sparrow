package main

import (
	"fmt"
	"os"
	"os/signal"
	"rpc-demo/center_server/base"
)
func close(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("quit,Got signal:", s)
}
func main()  {
	base.Init()
	if base.Start()==false{
		fmt.Println("start fail")
	}
	base.Run()
	close()
	base.Close()
}
