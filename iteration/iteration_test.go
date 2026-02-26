package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestStringsLibrary(t *testing.T) {
	t.Run("built-in repeat", func(t *testing.T) {
		got := strings.Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("check for substring", func(t *testing.T) {
		got := strings.Contains("hello world", "world")
		if got != true {
			t.Errorf("expected true but got %v", got)
		}
	})
}

// ExampleRepeat documents hot to use the function in godoc
func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
