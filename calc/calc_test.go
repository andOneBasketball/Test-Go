package main
 
import "testing"
 
//必须是Test开头
func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		//类似fmt.Printf
		t.Fatalf("add(2,4) err,expect:%d,actual:%d", 6, r)
	}
	t.Logf("test is success")
}
 
func TestSub(t *testing.T) {
	r := sub(10, 4)
	if r != 6 {
		//类似fmt.Printf
		t.Fatalf("add(2,4) err,expect:%d,actual:%d", 6, r)
	}
	t.Logf("test is success")
}