package sort_go
func InsertSort (num []int){
	n :=len(num)
	if n<=1 {
		return
	}
	for i:=1;i<n;i++ {
		key:=num[i]
		j:=i-1
		for j >= 0 && num[j] < key {
			num[j + 1] = num[j]
			j--
		}
		num[j + 1] = key
	}
}
