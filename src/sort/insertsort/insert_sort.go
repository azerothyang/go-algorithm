package main

import (
	"generator"
	"log"
	"time"
	"assert"
	"strconv"
)

func main() {
	arr := generator.GenerateArray(100000)
	startTime := time.Now().UnixNano()
	insertSort(arr)
	endTime := time.Now().UnixNano()
	if assert.Array(arr) == false {
		log.Println("算法有误")
	}
	dur := strconv.Itoa(int(endTime - startTime)/1000000)
	log.Println("总共用时:" + dur + " ms")
}

//插入排序
func insertSort(arr []int){
	//数组长度
	arrLength := len(arr)
	for i:=1; i<arrLength; i++{
		//和数组之前元素进行比较如果小于前一个元素则交换位置，否则跳出子循环
		for j:=i; j>0; j-- {
			if arr[j]<arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j] //交换两个数
			}else {
				break
			}
		}
	}
}
