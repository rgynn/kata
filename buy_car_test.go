package kata

import (
	"reflect"
	"testing"
)

func TestNbMonths(t *testing.T) {
	if result := NbMonths(2000, 8000, 1000, 1.5); !reflect.DeepEqual(result, [2]int{6, 766}) {
		t.Fatalf("expected result to be: []int{6, 766}, got: %+v", result)
	}
	if result := NbMonths(12000, 8000, 1000, 1.5); !reflect.DeepEqual(result, [2]int{0, 4000}) {
		t.Fatalf("expected result to be: []int{0, 4000}, got: %+v", result)
	}
	if result := NbMonths(8000, 12000, 500, 1.0); !reflect.DeepEqual(result, [2]int{8, 597}) {
		t.Fatalf("expected result to be: []int{8, 597}, got: %+v", result)
	}
	if result := NbMonths(18000, 32000, 1500, 1.25); !reflect.DeepEqual(result, [2]int{8, 332}) {
		t.Fatalf("expected result to be: []int{8, 332}, got: %+v", result)
	}
	if result := NbMonths(7500, 32000, 300, 1.55); !reflect.DeepEqual(result, [2]int{25, 122}) {
		t.Fatalf("expected result to be: []int{25, 122}, got: %+v", result)
	}
}
