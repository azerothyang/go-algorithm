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
	if assert.Array(arr) == false {
		log.Println("算法有误")
	}
	endTime := time.Now().UnixNano()
	dur := strconv.Itoa(int(endTime - startTime)/1000000)
	log.Println("总共用时:" + dur + " ms")
}

//插入排序
func insertSort(arr []int){
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
}
