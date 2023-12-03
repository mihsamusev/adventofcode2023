package main

import "testing"

func TestSingleDigit(t *testing.T) {
	actual := ParseNumberRefs("....*4....")
	if len(actual) == 0 {
		t.Errorf("bad boi")
	}
}

func TestMoreDigits(t *testing.T) {
	actual := ParseNumberRefs("....*44...")
	if len(actual) == 0 {
		t.Errorf("bad boi")
	}
}
