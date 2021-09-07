package kata

import (
	"reflect"
	"strings"
	"testing"
)

func TestMovingShift(t *testing.T) {
	input1 := "S mkx bocod"
	res1 := []string{"T p", "oc ", "iwl", "yo", ""}
	if result := MovingShift(input1, 1); !reflect.DeepEqual(result, res1) {
		t.Errorf("expected: ['%s'], got: ['%s']", strings.Join(res1, `','`), strings.Join(result, `','`))
	}

	input2 := "I should have known that you would have a perfect answer for me!!!"
	res2 := []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
	if result := MovingShift(input2, 1); !reflect.DeepEqual(result, res2) {
		t.Errorf(`expected: [%s], got: [%s]`, strings.Join(res2, `','`), strings.Join(result, `','`))
	}
}

func TestDemovingShift(t *testing.T) {

	input := []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
	res := "I should have known that you would have a perfect answer for me!!!"
	if result := DemovingShift(input, 1); result != res {
		t.Errorf(`expected: %s, got: %s`, res, result)
	}
}
