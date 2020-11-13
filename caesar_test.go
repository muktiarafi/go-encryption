package main

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	pt := "BELAJAR"
	k := 9
	want := "KNUJSJA"

	if got := CaesarEncrypt(pt, k); got != want {
		t.Errorf("Expected %q got %q instead", want, got)
	}
}
