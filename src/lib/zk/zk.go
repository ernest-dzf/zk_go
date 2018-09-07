package zk

import (
	samuelzk "github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"errors"
	. "loger"
	"util"
)
const (
	ZK_TIMEOUT	=	15
)
type ZkClient struct {
	samuelzk.Conn
	zkHost		[]string
	timeout		time.Duration
	conn *samuelzk.Conn
}
func (this *ZkClient) setZkHost(hosts []string) (){
	this.zkHost = hosts
}
func (this *ZkClient) setTimeout(duration time.Duration) () {
	this.timeout = duration
}

func NewZkClient(hosts []string, timeout time.Duration) (*ZkClient) {
	if len(hosts) == 0 {
		return nil
	}
	if timeout == 0 {
		timeout = ZK_TIMEOUT
	}
	client := &ZkClient{}
	for _, host := range hosts {
		if util.ValidIP4Address(host) == false && util.ValidServerAddress(host) == false {
			Error("parameter hosts illegal")
			return nil
		}
	}
	client.setTimeout(timeout)
	client.setZkHost(hosts)
	var err error
	client.conn, err = client.getZkConn()
	if err != nil {
		Error("getZkConn failed, " + err.Error())
		return nil
	}
	return client
}

//获取zk client 的连接
func (this *ZkClient) getZkConn() (*samuelzk.Conn, error) {
	if this.conn != nil {
		return this.conn, nil
	}
	newConn, _, err := samuelzk.Connect(this.zkHost, time.Second * this.timeout)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can not connect to zk server")
	}
	this.conn  = newConn
	return newConn, nil
}

//func (this *ZkClient)
