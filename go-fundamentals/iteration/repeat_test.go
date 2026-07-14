package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	//... setup ...
	for b.Loop() {
		// code to mesure ...
		Repeat("a")
	}
	// cleanup
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
