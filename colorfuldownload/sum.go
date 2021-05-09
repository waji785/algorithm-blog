package main



func Sum2(numbers []int,target int) bool{
	for i:=0;i<len(numbers);i++{
		for _,val:=range numbers{
			if target-numbers[i]==val||target-numbers[i]==0{
				return true
			}else if (target-numbers[i])>0{
				return Sum2(numbers,target-numbers[i])
			}
		}
	}
	return false
}