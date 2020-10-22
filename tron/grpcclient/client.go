package grpcclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type _conn struct {
	c          *grpc.ClientConn
	serverAddr string
}


func (c *_conn) Connect() (err error) {

	c.c, err = getClient(c.serverAddr)
	if nil != err {
		return err
	}
	// fmt.Printf("Connection status:%v\n", c.c.GetState())

	return
}

func (c *_conn) Close() {
	if nil != c {
		pushClient(c.serverAddr, c.c)
	}
}

func (c *_conn) ConnectOld() (err error) {

	c.c, err = grpc.Dial(c.serverAddr, grpc.WithInsecure())
	if nil != err {
		return err
	}
	// fmt.Printf("Connection status:%v\n", c.c.GetState())

	return
}

func (c *_conn) GetState() connectivity.State {
	return c.c.GetState()
}

func (c *_conn) Target() string {
	return c.c.Target()
}

func (c *_conn) CloseOld() {
	if nil != c {
		c.c.Close()
	}
}
