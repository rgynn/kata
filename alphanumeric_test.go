package kata

import "testing"

func TestAlphaNumeric(t *testing.T) {
	if result := alphanumeric(".*?"); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("a"); result != true {
		t.Fatalf("expected true, got %t", result)
	}
	if result := alphanumeric("Mazinkaiser"); result != true {
		t.Fatalf("expected true, got %t", result)
	}
	if result := alphanumeric("hello world_"); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("PassW0rd"); result != true {
		t.Fatalf("expected true, got %t", result)
	}
	if result := alphanumeric("     "); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric(""); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("\n\t\n"); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("ciao\n$$_"); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("__ * __"); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("&)))((("); result != false {
		t.Fatalf("expected false, got %t", result)
	}
	if result := alphanumeric("43534h56jmTHHF3k"); result != true {
		t.Fatalf("expected true, got %t", result)
	}
}
