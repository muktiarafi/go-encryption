package main

import "testing"

func TestVigenereEncryptUpper(t *testing.T) {
	pt := "BELAJAR"
	k := "SMART"
	want := "TQLRCSD"

	if got := VigenereEncrypt(pt, k); got != want {
		t.Errorf("Expected %q got %q instead", want, got)
	}
}
