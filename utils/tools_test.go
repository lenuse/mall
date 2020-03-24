package utils

import "testing"

func BenchmarkSlice2String(b *testing.B) {
	info := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		_, _ = Slice2String(info)
	}
}

func TestSlice2String(t *testing.T) {
	info := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := "1,2,3,4,5,6,7,8,9"
	str, _ := Slice2String(info)
	if str != result {
		t.Errorf("程序错误,%s", str)
	}
}
