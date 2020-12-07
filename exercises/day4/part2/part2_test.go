package main

import "testing"

// TestBirthYear Do a thing
func TestParseBirthYear(t *testing.T) {
	// PASS
	_, err := parseBirthYear("1920")
	if err != nil {
		t.Errorf("1920 should work")
	}
	_, err = parseBirthYear("2002")
	if err != nil {
		t.Errorf("2002 should work")
	}
	// FAILS
	_, err = parseBirthYear("1919")
	if err == nil {
		t.Errorf("1919 should fail")
	}
	_, err = parseBirthYear("2003")
	if err == nil {
		t.Errorf("2003 should fail")
	}
	_, err = parseBirthYear("abcd")
	if err == nil {
		t.Errorf("abcd should fail")
	}
}
