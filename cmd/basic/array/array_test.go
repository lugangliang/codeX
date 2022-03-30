package main

import (
	"codeX/util"
	"testing"
)

// 两个空数组
func TestMergeTwoSortedArray1(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray(nil, nil)
	if _, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error")
	}

}

// a2 为空
func TestMergeTwoSortedArray2(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray([]int{1,3,4}, nil)
	if _, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error")
	}
}

// a1 为空
func TestMergeTwoSortedArray3(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray(nil, []int{1,3,5})
	if _, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error")
	}
}

// a1,a2等长
func TestMergeTwoSortedArray4(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray([]int{4,8,10}, []int{1,3,40})
	if str, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error" + str)
	}
}

// a1 长度 > a2 长度
func TestMergeTwoSortedArray5(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray([]int{1,2,3,10, 60}, []int{1,3,40})
	if _, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error")
	}
}

// a1 长度 < a2 长度
func TestMergeTwoSortedArray6(t *testing.T) {
	var newArray []int

	newArray = MergeTwoSortedArray([]int{1,3,40}, []int{1,2,3,10, 60})
	if _, ok := util.CheckArraySortResult(newArray); !ok {
		t.Errorf("sort error")
	}
}