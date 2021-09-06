package kata

import "testing"

func TestNbDig(t *testing.T) {
	if result := NbDig(550, 5); result != 213 {
		t.Fatalf("expected 213 from 550, 5")
	}
}
