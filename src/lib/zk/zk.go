package zk

import (
	samuelzk "github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"errors"
	. "log"
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
	client := &ZkClient{}
	client.setTimeout(timeout)
	client.setZkHost(hosts)
	client.conn = nil
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
