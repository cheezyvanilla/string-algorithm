package app

import "testing"

func TestPalindromeCheck(t *testing.T){
	testCase := []struct{
		input string
		want bool
	}{
		{
			input: "kodok",
			want: true,
		},
		{
			input: "世界",
			want: false,
		},
		{
			input: "なるとをなると",
			want: false,
		},
	}

	for _, v := range testCase {
		got := PalindromeCheck(v.input)

		if got != v.want {
			t.Errorf("want %v got %v", v.want, got)
		}
	}
}