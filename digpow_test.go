package kata

import "testing"

func TestDigPow(t *testing.T) {
	if result := DigPow(89, 1); result != 1 {
		t.Fatalf("expected result to be: 1, got: %d", result)
	}
	if result := DigPow(92, 1); result != -1 {
		t.Fatalf("expected result to be: -1, got: %d", result)
	}
	if result := DigPow(3263, 1); result != -1 {
		t.Fatalf("expected result to be: -1, got: %d", result)
	}
}
