package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	got := ReverseString("testing123_#$!#")
	want := "#!$#_321gnitset"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
