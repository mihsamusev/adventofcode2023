package main

import "testing"

func TestBasic(t *testing.T) {
	digits := lineDigits("pqr3stu8vwx")
	value := calibrationValue(digits)
	expected := 38
	if value != expected {
		t.Errorf("value: %v != expected: %v", value, expected)
	}
}

func TestExtended(t *testing.T) {
	strs := []string{"pqr3stu8vwx", "xtwone3four"}
	expected := []int{38, 24}
	for i, str := range strs {
		digits := lineDigitsExtended(str)
		value := calibrationValue(digits)
		if value != expected[i] {
			t.Errorf("value: %v != expected: %v", value, expected)
		}
	}
}

func TestDigitsCorrect(t *testing.T) {
	str := "21xfxfourmzmqbqp1"
	actual := lineDigitsExtended(str)
	expected := []int{2, 1, 4, 1}
	AssertSlicesEqual(t, actual, expected)
}

func TestDigitsCorrectOneight(t *testing.T) {
	str := "oneightwoneight"
	actual := lineDigitsExtended(str)
	expected := []int{1, 8, 2, 1, 8}
	AssertSlicesEqual(t, actual, expected)
}

func TestDigitsCorrectTwone(t *testing.T) {
	str := "twone"
	actual := lineDigitsExtended(str)
	expected := []int{2, 1}
	AssertSlicesEqual(t, actual, expected)
}

func AssertSlicesEqual(t *testing.T, actual, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("actual: %v != expected: %v", actual, expected)
		return
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("actual: %v != expected: %v", actual, expected)
			return
		}
	}
}
