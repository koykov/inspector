package inspector

import "testing"

func TestParseHelpers(t *testing.T) {
	if v, ok := atoi("42"); !ok || v != 42 {
		t.Fatalf("atoi: %d %v", v, ok)
	}
	if _, ok := atoi("x"); ok {
		t.Fatal("atoi must fail")
	}
	if _, ok := atoi("999999999999999999999"); ok {
		t.Fatal("atoi overflow must fail")
	}
	if v, ok := atou("42"); !ok || v != 42 {
		t.Fatalf("atou: %d %v", v, ok)
	}
	if _, ok := atou("-1"); ok {
		t.Fatal("atou must reject negative")
	}
	if v, ok := atof("3.5"); !ok || v != 3.5 {
		t.Fatalf("atof: %v %v", v, ok)
	}
	if _, ok := atof("x"); ok {
		t.Fatal("atof must fail")
	}
}
