package main
//TCP并发端口扫描器
import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main1(){
	start:=time.Now()
	var wg sync.WaitGroup
	for i:=1;i<65535;i++{
		go func(j int){
			defer wg.Done()
			wg.Add(1)
			address :=fmt.Sprintf("192.168.0.1:%d",j)
			conn,err :=net.Dial("tcp",address)
			if err !=nil{
				fmt.Println("closed...",address)
				return
			}
			conn.Close()
			fmt.Printf("opened",address)
		}(i)
	}
	wg.Wait()
	elapsed:=time.Since(start)/1e9
	fmt.Printf("\n\n%d seconds",elapsed)
}
