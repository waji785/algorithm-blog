package sort_go

func BubbleSort(num []int){
	n :=len(num)
	if n<=1 {
		return
	}
	for i:=0;i<n;i++ {
		for j:=i+1;j<n;j++ {
			if (num)[i]>(num)[j]{
				(num)[i],(num)[j]=(num)[j],(num)[i]
			}
		}
	}
}