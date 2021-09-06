package kata

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	if result := MaximumSubarraySum([]int{}); result != 0 {
		t.Fatalf("expeted 0, got: %d", result)
	}
	if result := MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); result != 6 {
		t.Fatalf("expeted 6, got: %d", result)
	}
	if result := MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}); result != 0 {
		t.Fatalf("expeted 6, got: %d", result)
	}
}
