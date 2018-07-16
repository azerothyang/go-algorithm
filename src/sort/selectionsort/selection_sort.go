package main

import (
	"generator"
	"time"
	"assert"
	"log"
	"strconv"
)

func main() {
	arr := generator.GenerateArray(100000)
	startTime := time.Now().UnixNano()
	selectionSort(arr)
	endTime := time.Now().UnixNano()
	if assert.Array(arr) == false {
		log.Println("算法有误")
	}
	dur := strconv.Itoa(int(endTime - startTime)/1000000)
	log.Println("总共用时:" + dur + " ms")
}

//选择排序
func selectionSort(arr []int)  {
	len := len(arr)
	for i:=0; i<len; i++{
		tmp := arr[i]
		key := i
		for j:=i+1; j<len; j++{
			if arr[j] < tmp {
				tmp = arr[j]
				key = j
			}
		}
		arr[i], arr[key] = arr[key], arr[i]
	}
}


