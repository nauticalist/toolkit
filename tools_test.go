package toolkit

import "testing"

func TestTools_RandromString(t * testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)

	if len(s) != 10 {
		t.Error("wrong lenght random string returned")
	}
}