package app

import "testing"

func TestVowelConsonantCounter(t *testing.T){
	t.Run("with int", func(t *testing.T){
		gotVow, gotCons := VowelConsonantCounter("awkokwo69")
		wantVow, wantCons := 3, 4

		if gotVow != wantVow || gotCons != wantCons {
			t.Errorf("want %d, %d got %d, %d", gotVow, gotCons, wantVow, wantCons)
		}
	})

	t.Run("with spaces", func(t *testing.T) {
		gotVow, gotCons := VowelConsonantCounter("aku adalah alien")
		wantVow, wantCons := 8, 6

		if gotVow != wantVow || gotCons != wantCons {
			t.Errorf("want %d, %d got %d, %d", wantVow, wantCons, gotVow, gotVow)
		}
	})
}