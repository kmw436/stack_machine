package main

import (
	"testing"
)

func TestWhoAmI(t *testing.T){
	result := whoAmI()
	if result != "Stack Machine" {
		t.Errorf("I think I am a %q", result)
	}
}