package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("b", 5)
	fmt.Println(got)
	// Output: bbbbb
}
