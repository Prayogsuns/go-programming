package word

import (
	"fmt"
	"strings"
	"testing"
)

var ts string = "I was very old when I was a child. My World came crashing when I was young. When I was older I never had the wisdom. I was elder then I was conscious."

func ExampleCount() {
	WordCount := Count("I was very old when I was a child. My World came crashing when I was young. When I was older I never had the wisdom. I was elder then I was conscious.")
	fmt.Println(WordCount)
	// Output:
	// 33
}


func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(ts)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(ts)
	}
}

func TestUseCount(t *testing.T) {
	tss := strings.Fields(ts)
	
	tsm := make(map[string]int)
	for _, v := range tss {
		tsm[v]++
	}	
	// fmt.Println("Map", tsm)
	
	testVar := true
	wc := UseCount(ts)
	if wc != nil {
		if len(wc) == len(tsm) {
			for k, v := range tsm {
				if tv, ok := wc[k]; !ok || tv != v {
					testVar = false
				}
			}
		} else {
			testVar = false
		}
	} else {
		testVar = false
	}
	if !testVar {
		t.Errorf("Expected %#v, %s %#v", tsm, "Got", wc)
	}
}

func TestCount(t *testing.T) {
	lentss := Count(ts)

	if lentss != 33 {
		t.Errorf("Expected %d, %s, %v", 33, "Got", lentss)	
	}	
}
