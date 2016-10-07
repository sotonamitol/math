package main

import "testing"

func TestF(t *testing.T) {
	if f(4) != 24 {
		t.Errorf("err")
	}
	if f(5) != 120 {
		t.Errorf("err")
	}
	if f(6) != 721 {
		t.Errorf("err")
	}
}
