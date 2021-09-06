package kata

import "testing"

func TestDigitalRoot(t *testing.T) {
	if result := DigitalRoot(0); result != 0 {
		t.Fatalf("expected 0 from 0")
	}
	if result := DigitalRoot(10); result != 1 {
		t.Fatalf("expected 1 from 10")
	}
	if result := DigitalRoot(16); result != 7 {
		t.Fatalf("expected 7 from 16")
	}
	if result := DigitalRoot(195); result != 6 {
		t.Fatalf("expected 6 from 195, got: %d", result)
	}
	if result := DigitalRoot(992); result != 2 {
		t.Fatalf("expected 2 from 992, got: %d", result)
	}
	if result := DigitalRoot(167346); result != 9 {
		t.Fatalf("expected 9 from 167346, got: %d", result)
	}
}
