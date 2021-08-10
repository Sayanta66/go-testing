package main

import "testing"

func Test_counter(t *testing.T)  {
	start, add := 10, 10
	want := start + add
	got := Count(start, add)

	if want != got {
		t.Fatalf("wanted %d, got %d", want, got)
	}

}