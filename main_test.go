package main

import "testing"

func Test_add(t *testing.T)  {
	start, plus := 10, 10
	want := start + plus
	got := add(start, plus)

	if want != got {
		t.Fatalf("wanted %d, got %d", want, got)
	}

}

func Test_multiply(t *testing.T) {
	start, times:= 10, 10
	want := start * times
	got := multiply(start, times)

	if want != got {
		t.Fatalf("wanted %d, got %d", want, got)
	}

}