package main

import "fmt"
var arr = [10]int{67,7,434,3,45,678,74,463,32,3}
func BubbleSort(num *[10]int){
	n :=len(*num)
	if n<=1 {
		return
	}
	for i:=0;i<n;i++ {
		for j:=i+1;j<n;j++ {
			if (*num)[i]>(*num)[j]{
				(*num)[i],(*num)[j]=(*num)[j],(*num)[i]
			}
		}
	}
}

func main()  {
	BubbleSort(&arr)
	fmt.Printf("%v",arr)
}