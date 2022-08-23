package tooklit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandmoString(10)
	if len(s) != 10 {
		t.Error("wrong length")
	}

}
