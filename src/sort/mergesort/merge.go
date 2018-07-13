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
	arr = mergeSort(arr)
	if assert.Array(arr) == false {
		log.Println("算法有误")
	}
	endTime := time.Now().UnixNano()
	dur := strconv.Itoa(int(endTime - startTime)/1000000)
	log.Println("总共用时:" + dur + " ms")
}


//归并排序, 从小到大
func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2 //切成两段
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

//合并排好序的两个数组
func merge(left, right []int) []int {
	lLen, rLen := len(left), len(right)
	result := make([]int, 0)
	l, r := 0, 0
	for l < lLen && r < rLen {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	if l < lLen {
		result = append(result, left[l:]...)
	} else if r < rLen {
		result = append(result, right[r:]...)
	}
	return result
}
