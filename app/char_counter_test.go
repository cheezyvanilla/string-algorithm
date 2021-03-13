package app

import (
	"reflect"
	"testing"
)

func TestCharCounter(t *testing.T){
	got := CharCounter("ababababab")
	want := map[string]int{
		"a": 5,
		"b": 5,
	}

	if !reflect.DeepEqual(got, want){
		t.Errorf("want %d, %d got %d, %d", want["a"], want["b"], got["a"], got["b"])
	}
}