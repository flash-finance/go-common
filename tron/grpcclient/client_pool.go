package grpcclient

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"strings"
	"sync"
	"time"
)


var clientPool sync.Map

var _cleanOnce sync.Once


var (
	DefaultPoolSize    = 300
	DefaultTimeoutTime = 300 * time.Second
)


type IdleClientType struct {
	c *grpc.ClientConn
	t time.Time
}


var (
	ErrTargetNotMatch     = fmt.Errorf("client target not match")
	ErrClientChanFull     = fmt.Errorf("client chan full")
	ErrInvalidClientState = fmt.Errorf("client state invalid")
)


func getClient(target string) (*grpc.ClientConn, error) {
	c, ok := clientPool.Load(target)

	var ccb *ClientConnBuffer
	if ok && nil != c {
		ccb, ok = c.(*ClientConnBuffer)
		if !ok || nil == ccb {
			ccb = NewClientConnBuffer(target, DefaultPoolSize)
			clientPool.Store(target, ccb)
		}
	} else {
		ccb = NewClientConnBuffer(target, DefaultPoolSize)
		clientPool.Store(target, ccb)
	}

	return ccb.GetConnection()
}


func pushClient(target string, c *grpc.ClientConn) error {
	cc, ok := clientPool.Load(target)

	var ccb *ClientConnBuffer
	if ok && nil != cc {
		ccb, ok = cc.(*ClientConnBuffer)
		if !ok || nil == ccb {
			ccb = NewClientConnBuffer(target, DefaultPoolSize)
		}
	} else {
		ccb = NewClientConnBuffer(target, DefaultPoolSize)
		clientPool.Store(target, ccb)
	}
	return ccb.PushConnection(c)
}


func GetConnectionPoolState() string {

	ret := make([]string, 0)
	clientPool.Range(func(key, val interface{}) bool {
		ccb, ok := val.(*ClientConnBuffer)
		if ok && nil != ccb {
			ret = append(ret, fmt.Sprintf("target:[%v], connection pool current size:%v",
				key, len(ccb.ClientChan)))
		}
		return true
	})
	return strings.Join(ret, "\n")
}


func NewClientConnBuffer(target string, poolSize int) *ClientConnBuffer {
	_cleanOnce.Do(func() {
		go cleanConnectionPoolTask()
	})
	// fmt.Printf("NewClientConnBuffer for target:[%v], poolSize:[%v]\n", target, poolSize)
	return &ClientConnBuffer{
		Target:     target,
		ClientChan: make(chan *IdleClientType, poolSize),
		PoolSize:   poolSize,
	}
}


type ClientConnBuffer struct {
	Target     string
	ClientChan chan *IdleClientType
	PoolSize   int
}

func (ccb *ClientConnBuffer) GetConnection() (ret *grpc.ClientConn, err error) {
	select {
	case retR := <-ccb.ClientChan:
		ret = retR.c
	default:
		ret, err = grpc.Dial(ccb.Target, grpc.WithInsecure())
		if nil != err {
			return nil, err
		}
	}
	return ret, err
}

func (ccb *ClientConnBuffer) PushConnection(c *grpc.ClientConn) (err error) {
	if nil == c {
		// fmt.Printf("Close Target err:[%v] grpcclient connection for target:[%v]\n", c.Target(), ccb.Target)
		return nil
	}
	if c.Target() != ccb.Target {
		c.Close()
		return ErrTargetNotMatch
	}
	stat := c.GetState()
	if connectivity.Idle == stat || connectivity.Ready == stat {
		select {
		case ccb.ClientChan <- &IdleClientType{
			c: c,
			t: time.Now(),
		}:
			return nil
		default:
			c.Close()
			return ErrClientChanFull
		}
	}
	c.Close()
	return ErrInvalidClientState
}

func cleanConnectionPoolTask() {
	for {
		time.Sleep(DefaultTimeoutTime)
		CleanConnectionPool()
	}
}

func CleanConnectionPool() {

	timeoutCnt := 0
	restCnt := 0
	clientPool.Range(func(key, val interface{}) bool {

		ccb, ok := val.(*ClientConnBuffer)
		if ok && nil != ccb {

			cnt := len(ccb.ClientChan)
			for i := 0; i < cnt; i++ {
				idc := <-ccb.ClientChan
				if time.Since(idc.t) > DefaultTimeoutTime {
					idc.c.Close()
					timeoutCnt++
				} else {
					restCnt++
					ccb.ClientChan <- idc
				}
			}
		}
		return true
	})

	fmt.Printf("[%v] clean timeout connection count:[%v], left:[%v]\n",
		time.Now().Format("2006-01-02 15:04:05.000"), timeoutCnt, restCnt)
	fmt.Println(GetConnectionPoolState())
}
