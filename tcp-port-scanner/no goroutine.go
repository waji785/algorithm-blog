package main
//无并发版本
import (
	"fmt"
	"net"
)

func main2(){
	for i:=1;i<120;i++{
		address :=fmt.Sprintf("192.168.0.1:%d",i)
		conn,err :=net.Dial("tcp",address)
		if err !=nil{
			fmt.Println("closed...",address)
			continue
		}
		conn.Close()
		fmt.Printf("opened",address)
	}
}