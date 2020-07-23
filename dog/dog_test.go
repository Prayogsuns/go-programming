package doggy

import (
        "fmt"
        "testing"
)

func TestYears(t *testing.T) {
        got := Years(2)
        if got != 14 {
                t.Errorf("Expected Years(2) == %d, but %s %d\n", 14, "got", got)
        }
}

func TestYearsTwo(t *testing.T) {
        got := YearsTwo(3)
        if got != 21 {
                t.Errorf("Expected YearsTwo(3) == %d, but %s %d\n", 21, "got", got)
        }
}

func ExampleYears() {
	got := Years(4)
        fmt.Println(got)
        // output:
        // 28
}

func ExampleYearsTwo() {
	got := YearsTwo(5)
        fmt.Println(got)
        // output:
        // 35
}

func BenchmarkYears(b *testing.B) {
        for i := 0; i < b.N; i++ {
                Years(9)
        }
}

func BenchmarkYearsTwo(b *testing.B) {
        for i := 0; i < b.N; i++ {
                YearsTwo(9)
        }
}
