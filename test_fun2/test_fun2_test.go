package main

// go test fun2_test.go

import "testing"  


func TestFun2(t *testing.T) {
	a, b := f2()
	if a != 3 && b != false {
		t.Errorf("f2() = (%d, %t); want (3, false)", a, b)
	}
}
