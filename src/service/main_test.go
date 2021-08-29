package main

import (
	"testing"
)

func TestMain(t *testing) {
	expected := 1
	got := 1
	if expected != got {
		t.Errorf("main test failed: got %v wanted %v", got, expected)
	}
}