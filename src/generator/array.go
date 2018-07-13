package generator

import "math/rand"

//生成一个含有n个元素的切片
func GenerateArray(n int)[]int {
	s := make([]int, n)
	for i:=0; i<n; i++ {
		s[i] = rand.Int()
	}
	return s
}

