package main

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	pt := "BELAJAR"
	k := 9

	if c := CaesarEncrypt(pt, k); c != "KNUJSJA" {
		t.Errorf("Expected KNUJSJA but got %s", c)
	}
}
