package main

import "testing"

func TestGetMiddle(t *testing.T) {
	if test := GetMiddle("hello"); test != "l" {
		t.Fatalf(`GetMiddle("hello") should be %v but was %v`, "l", test)
	}
	if test := GetMiddle("wuta"); test != "ut" {
		t.Fatalf(`GetMiddle("wuta") should be %v but was %v`, "ut", test)
	}
	if test := GetMiddle("hellothissourcery"); test != "s" {
		t.Fatalf(`GetMiddle("hellothissourcery") should be %v but was %v`, "s", test)
	}
	if test := GetMiddle("hellothissourcery"); test != "s" {
		t.Fatalf(`GetMiddle("hellothissourcery") should be %v but was %v`, "s", test)
	}
}
