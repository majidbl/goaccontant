package model

import "testing"

func TestConnectedDB(t *testing.T) {
	err := Dbcheck()
	if err != nil {
		t.Errorf("expected nil but %s occurred!!", err)
	}
}
