package main
import (
	"fmt"
	. "loger"
	"util"
)
const (
	ZK_HOST1 = "134.175.195.174:2181"
)

type TTTT struct {
	o string
	OO	string
}
func (this *TTTT) test() {
	fmt.Println(this.o)
}
func main() {
	fmt.Println(util.ValidServerAddress("192.168.19.1:65536"))
	Info("dfdfdfd")
	Error(197865,"test")
	Error("48239843")
	//t := &TTTT{}
	//t.o = "ssss"
	//t.test()
	//fmt.Printf("%v \n", children)
	//children, stat, ch, err := c.ChildrenW("/")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v %+v\n", children, stat)
	//e := <-ch
	//fmt.Printf("%+v\n", e)
}
