package leader

import (
	"rpc-demo/common/frame"
	"rpc-demo/common/frame/rpc_protocol"
	"rpc-demo/common/rpc_service"
)

type leader struct {
	frame.Frame
	rpcClientMap map[uint32]*rpc_service.RpcClient
	serverList []rpc_protocol.ServiceItem
}

func (this* leader) Init()  {
	this.Frame.Init("center_config.xml")
	//rpc

}
func (this* leader) Start() bool {
	this.Frame.Start()

	return true
}

func (this* leader) Run() error {
	this.Frame.Run()
	return nil
}

func (this* leader) Close() {
	this.Frame.Close()
}