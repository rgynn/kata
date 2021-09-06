package kata

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {
	if result := MultiplicationTable(1); !reflect.DeepEqual(result, [][]int{{1}}) {
		t.Fatalf("expected result to be: [][]int{{1}}, got: %v", result)
	}
	if result := MultiplicationTable(2); !reflect.DeepEqual(result, [][]int{{1, 2}, {2, 4}}) {
		t.Fatalf("expected result to be: [][]int{{1}}, got: %v", result)
	}
	if result := MultiplicationTable(3); !reflect.DeepEqual(result, [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}) {
		t.Fatalf("expected result to be: [][]int{{1}}, got: %v", result)
	}
}
