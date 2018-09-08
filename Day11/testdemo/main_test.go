package main

import "testing"

func TestAdd(t *testing.T) {
	i := add(1, 3)
	if i != 4{
		t.Fatalf("测试失败")
	}else {
		t.Log("success test")
	}
}
