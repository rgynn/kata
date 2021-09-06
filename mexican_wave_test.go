package kata

import (
	"reflect"
	"testing"
)

func TestMexicanWave(t *testing.T) {
	if result := wave(" x yz"); !reflect.DeepEqual(result, []string{" X yz", " x Yz", " x yZ"}) {
		t.Fatalf("expected ' x yz' to equal: [' X yz', ' x Yz', ' x yZ']")
	}
}
