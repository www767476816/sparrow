package client

import "sparrow/dispatch_server/connect"

type Client struct {
	conn *connect.Connect
}

func CreateClient(c *connect.Connect) *Client{
	result:=new(Client)
	result.conn=c
	return result
}

func (this* Client) Dispatch(data []byte)  {

}