package mysum

import "testing"

func TestMySum(t *testing.T) {
	got := MySum(1,9,19)
	if got != 29 {
		t.Errorf("MySum(1,9,19) = %d; Expected %s, %d", 29, "got", got)
	}
}