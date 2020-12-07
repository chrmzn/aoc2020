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

func TestParseIssueYear(t *testing.T) {
	// PASS
	_, err := parseIssueYear("2010")
	if err != nil {
		t.Errorf("2010 should work")
	}
	_, err = parseIssueYear("2020")
	if err != nil {
		t.Errorf("2020 should work")
	}
	// FAILS
	_, err = parseIssueYear("2009")
	if err == nil {
		t.Errorf("2009 should fail")
	}
	_, err = parseIssueYear("2021")
	if err == nil {
		t.Errorf("2021 should fail")
	}
	_, err = parseIssueYear("abcd")
	if err == nil {
		t.Errorf("abcd should fail")
	}
}
