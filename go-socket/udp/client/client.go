package main

import (
	"fmt"
	"net"
)

func main(){
	//建立连接
	listen,err:=net.DialUDP("udp",nil,&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port: 9090,
	})
	if err !=nil{
		fmt.Printf("listen udp server error:%v\n",err)
	}
	defer listen.Close()
	//发送数据
	sendData:=[]byte("hello")
	_,err=listen.Write(sendData)
	if err!=nil{
		fmt.Println("fail to send message:",err)
		return
	}
	//接收数据
	data:=make([]byte,4096)
	n,remoteAddr,err:=listen.ReadFromUDP(data)
	if err!=nil{
		fmt.Println("fail to receive message,err:",err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}