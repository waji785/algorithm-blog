package threadpoolgo

import "fmt"

func main(){
	pool:=NewPool(10,50)
	defer pool.Release()

	pool.WaitTotal(8)

	for i := 0; i < 8; i++ {
		count:=i
		pool.AddTask(func(){
			fmt.Printf("worker! Number %d\n",count)
			defer pool.Done()
		})
	}
	pool.Wait()
}
