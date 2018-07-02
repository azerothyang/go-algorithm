package main

import "fmt"

func main() {
	var arr = []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertSort(arr))
}

//插入排序
func insertSort(arr []int) []int{
	//数组长度
	arrLength := len(arr)
	for i:=1; i<arrLength; i++{
		tmp := arr[i]
		j := i - 1
		//和数组之前的做比较
		for j>=0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
	return arr
}
