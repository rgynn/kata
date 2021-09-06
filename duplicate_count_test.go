package kata

import "testing"

func TestDuplicateCount(t *testing.T) {
	if result := duplicate_count("abcde"); result != 0 {
		t.Fatalf("expected result to be 0, got: %d", result)
	}
	if result := duplicate_count("abcdeaB11"); result != 3 {
		t.Fatalf("expected result to be 3, got: %d", result)
	}
	if result := duplicate_count("indivisibility"); result != 1 {
		t.Fatalf("expected result to be 1, got: %d", result)
	}
}
