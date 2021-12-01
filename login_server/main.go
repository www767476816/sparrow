package main

import (
	"fmt"
	"os"
	"os/signal"
	"sparrow/common"
	"sparrow/login_server/base"
)
func close(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("quit,Got signal:", s)
}
func main()  {
	common.DuildMode=os.Args[1]

	base.Init()
	if base.Start()==false{
		fmt.Println("start fail")
	}
	base.Run()
	close()
	base.Close()
}
