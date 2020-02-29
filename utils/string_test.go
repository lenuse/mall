package utils

import "testing"

// 单元测试
func TestStringToUpper(t *testing.T) {
	str := "public_b2c_goods_brand"
	res := StringToUpper(str, '_')
	expected := "PublicB2cGoodsBrand"
	if res != expected {
		t.Errorf("toUpper:%s, result:%s",str, res)
	}
}

// 性能测试
func BenchmarkStringToUpper(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	str := "public_b2c_goods_brand"
	for i := 0; i < b.N; i++ {
		StringToUpper(str, '_')
	}
}

// 并发性能测试
func BenchmarkStringToUpperParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			str := "public_b2c_goods_brand"
			StringToUpper(str, '_')
		}
	})

}
