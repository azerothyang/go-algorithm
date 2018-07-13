package assert

//判断数组是否从小到大排好序
func Array(arr []int) bool{
	n := len(arr)
	for i:=0; i<n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
