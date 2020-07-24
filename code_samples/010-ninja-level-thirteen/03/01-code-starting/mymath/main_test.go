package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		in []int
		out float64
	}

	tests := []test{
		test{[]int{1,2,3,4,5,6,7,8,9,10,11},6.0},
		test{[]int{1,2,6,100},4.0},
		test{[]int{0,2,4,8,10,16,1000},8.0},
	}	

	for _, test := range tests {
		got := CenteredAvg(test.in)	
	
		if got != test.out {
			t.Errorf("Expected %f, %s %f", test.out, "Got", got)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1,2,3,4}))
	// Output: 2.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1,2,3,4,5,6,7,8,9,10})
	}
}