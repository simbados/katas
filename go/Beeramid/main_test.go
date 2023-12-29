package main

import (
	"testing"
)

func Test1(t *testing.T) {
	actual := Beeramid(1500, 2)
	expected := 12
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}

func Test2(t *testing.T) {
	actual := Beeramid(21, 1.5)
	expected := 3
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}

func Test3(t *testing.T) {
	actual := Beeramid(-1, 4.0)
	expected := 0
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}

func Test4(t *testing.T) {
	actual := Beeramid(5000, 3)
	expected := 16
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}

func Test5(t *testing.T) {
	actual := Beeramid(9, 2)
	expected := 1
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}

func Test6(t *testing.T) {
	actual := Beeramid(454, 5)
	expected := 5
	if actual != expected {
		t.Errorf("%d should equal %d", actual, expected)
	}
}
