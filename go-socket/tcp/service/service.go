package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()
	//当前连接的发送和接受操作
	for{
		reader:=bufio.NewReader(conn)
		var buf [128]byte
		n,err:=reader.Read(buf[:])
		if err!=nil{
			fmt.Printf("read conn failed,err:%v\n",err)
			break
		}
		recv:=string(buf[:n])
		fmt.Printf("received messsage:%v\n",recv)
		//将收到的数据返回给客户端
		_,err=conn.Write([]byte("ok"))
		if err !=nil{
			fmt.Printf("write from conn failed,err:%v\n",err)
			break
		}
	}
}
func main(){
	listen,err:=net.Listen("tcp","127.0.0.1:9090")
	if err !=nil{
		fmt.Printf("listen failed,err:%v\n",err)
		return
	}
	for{
		//等待客户端建立连接
		conn,err:=listen.Accept()
		if err !=nil{
			fmt.Printf("accepct failed,err:%v\n",err)
			continue
		}
		//启动一个单独的goroutine
		go process(conn)
	}
}