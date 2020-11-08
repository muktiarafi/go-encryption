package main

import "testing"

func TestVigenereEncryptUpper(t *testing.T) {
	pt := "BELAJAR"
	k := "SMART"

	if c := VigenereEncrypt(pt, k); c != "TQLRCSD" {
		t.Errorf("Expected TQLRCSD, got %s", c)
	}
}

func TestVigenereLower(t *testing.T) {
	pt := "belajar"
	k := "smart"

	if c := VigenereEncrypt(pt, k); c != "tqlrcsd" {
		t.Errorf("Expected tqlrcsd, got %s", c)
	}
}
