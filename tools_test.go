package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var tesTools Tools

	s := tesTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
}
