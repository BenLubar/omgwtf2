package main

import "testing"

func Testr(t *testing.T) {
	r1 := r()
	r2 := r()
	r3 := r()

	if r1 == r2 && r3 == r3 {
		t.Fail()
	}
}

func Testrr(t *testing.T) {
	rr1 := rr()
	rr2 := rr()
	rr3 := rr()

	if rr1 == rr2 && rr3 == rr3 {
		t.Fail()
	}
}

func Testrrr(t *testing.T) {
	rrr1 := rrr()
	rrr2 := rrr()
	rrr3 := rrr()

	if rrr1 == rrr2 && rrr3 == rrr3 {
		t.Fail()
	}
}
