package main

import (
	"testing"
)

// 使用方法 go test -test.bench=".*" -v

func BenchmarkMergeTwoSortedArray(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeTwoSortedArray([]int{3,4,8}, []int{1,12,40})
	}
}

func BenchmarkMergeTwoSortedArray1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeTwoSortedArray(nil, nil)
	}
}

func BenchmarkMergeTwoSortedArray2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeTwoSortedArray([]int{}, []int{})
	}
}