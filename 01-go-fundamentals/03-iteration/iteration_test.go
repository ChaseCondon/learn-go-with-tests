package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeat character 8 times", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertCorrectMessage(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeatedString := Repeat("a", 5)
	fmt.Println(repeatedString)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
