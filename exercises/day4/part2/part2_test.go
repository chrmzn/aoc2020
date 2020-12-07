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

func TestParseExpYear(t *testing.T) {
	// PASS
	_, err := parseExpYear("2020")
	if err != nil {
		t.Errorf("2020 should work")
	}
	_, err = parseExpYear("2030")
	if err != nil {
		t.Errorf("2030 should work")
	}
	// FAILS
	_, err = parseExpYear("2019")
	if err == nil {
		t.Errorf("2019 should fail")
	}
	_, err = parseExpYear("2031")
	if err == nil {
		t.Errorf("2031 should fail")
	}
	_, err = parseExpYear("abcd")
	if err == nil {
		t.Errorf("abcd should fail")
	}
}

func TestParseHeightStringCm(t *testing.T) {
	height, scale, err := parseHeightString("160cm")
	if err != nil {
		t.Errorf("Could not parse 160cm")
	}
	if height != 160 {
		t.Errorf("Height should be 160, got %d", height)
	}
	if scale != "cm" {
		t.Errorf("Failed to get cm scale for 160cm")
	}
	_, _, err = parseHeightString("149cm")
	if err == nil {
		t.Errorf("149cm is outside lower bound")
	}
	_, _, err = parseHeightString("194cm")
	if err == nil {
		t.Errorf("194cm is outside upper bound")
	}
}

func TestParseHeightStringIn(t *testing.T) {
	height, scale, err := parseHeightString("65in")
	if err != nil {
		t.Errorf("Could not parse 65in")
	}
	if height != 65 {
		t.Errorf("Height should be 65, got %d", height)
	}
	if scale != "in" {
		t.Errorf("Failed to get in scale for 65in, got %s", scale)
	}
	_, _, err = parseHeightString("58in")
	if err == nil {
		t.Errorf("58in is outside lower bound")
	}
	_, _, err = parseHeightString("77in")
	if err == nil {
		t.Errorf("77in is outside upper bound")
	}
}

func TestParseHeightStringFail(t *testing.T) {
	_, _, err := parseHeightString("160")
	if err == nil {
		t.Errorf("160 should fail")
	}
	_, _, err = parseHeightString("160")
	if err == nil {
		t.Errorf("160 should fail")
	}
}
