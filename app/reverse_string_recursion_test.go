package app

import "testing"

func TestReverseString(t *testing.T){
	got := ReverseString("anaconda")
	want := "adnocana"

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}