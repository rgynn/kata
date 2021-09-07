package kata

import (
	"reflect"
	"testing"
)

func TestPrimeGap(t *testing.T) {
	if result := Gap(2, 5, 5); result != nil {
		t.Fatalf("expected: nil, got: %v", result)
	}
	if result := Gap(2, 5, 7); !reflect.DeepEqual(result, []int{5, 7}) {
		t.Fatalf("expected: []int{5, 7}, got: %v", result)
	}
	if result := Gap(4, 130, 200); !reflect.DeepEqual(result, []int{163, 167}) {
		t.Fatalf("expected: []int{163, 167}, got: %v", result)
	}
	if result := Gap(2, 100, 110); !reflect.DeepEqual(result, []int{101, 103}) {
		t.Fatalf("expected: []int{101, 103}, got: %v", result)
	}
	if result := Gap(4, 100, 110); !reflect.DeepEqual(result, []int{103, 107}) {
		t.Fatalf("expected: []int{103, 107}, got: %v", result)
	}
	if result := Gap(6, 100, 110); result != nil {
		t.Fatalf("expected: nil, got: %v", result)
	}
	if result := Gap(8, 300, 400); !reflect.DeepEqual(result, []int{359, 367}) {
		t.Fatalf("expected: []int{359, 367}, got: %v", result)
	}
	if result := Gap(10, 300, 400); !reflect.DeepEqual(result, []int{337, 347}) {
		t.Fatalf("expected: []int{337, 347}, got: %v", result)
	}
}
