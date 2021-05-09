package main
//goroutine池版
import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int,result chan int){
	for p:=range  ports{
		address:=fmt.Sprint("192.168.0.1:%d",p)
		conn,err:=net.Dial("tcp",address)
		if err!=nil{
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}
func main(){
	ports :=make(chan int,100)
	result:=make(chan int)
	var openports []int
	var closedports []int
	for i:=0;i<cap(ports);i++{
		go worker(ports,result)
	}
	go func(){
		for i:=1;i<1024;i++{
			ports <- i
		}
	}()
	for i:=1;i<1024;i++{
		port := <-result
		if port !=0{
			openports=append(openports,port)
		}else{
			closedports=append(closedports,port)
		}
	}
	close(ports)
	close(result)
	sort.Ints(openports)
	sort.Ints(closedports)
	for _,port:=range closedports{
		fmt.Println(port)
	}
}