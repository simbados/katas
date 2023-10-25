package main

import "testing"

func TestNbYear(t *testing.T) {
	if test := NbYear(1500, 5, 100, 5000); test != 15 {
		t.Fatalf(`NbYear(1500, 5, 100, 5000) should be %v but was %v`, 15, test)
	}
	if test := NbYear(1500000, 2.5, 10000, 2000000); test != 10 {
		t.Fatalf(`NbYear(1500000, 2.5, 10000, 2000000) test should be %v but was %v`, 10, test)
	}
	if test := NbYear(1500000, 0.25, 1000, 2000000); test != 94 {
		t.Fatalf(`NbYear(1500000, 0.25, 1000, 2000000) test should be %v but was %v`, 94, test)
	}
	if test := NbYear(1500000, 0.25, -1000, 2000000); test != 151 {
		t.Fatalf(`NbYear(1500000, 0.25, -1000, 2000000) test should be %v but was %v`, 151, test)
	}
}
