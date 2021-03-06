// Copyright 2016-2017 The psh Authors. All rights reserved.
package psh

import (
	"os"
	"testing"
)

func TestSegmentVirtualEnvCompile(t *testing.T) {
	err := os.Setenv("VIRTUAL_ENV", "/home/user/venv")
	if err != nil {
		t.Fatalf("Can't set the required enviornment.")
	}
	expected := `venv`
	segment := NewSegmentVirtualEnv()
	segment.Compile()
	if string(segment.Data) != expected {
		t.Fatalf("Compiled data expected to be %q but got %q", expected, string(segment.Data))
	}
}

func TestSegmentVirtualEnvRender(t *testing.T) {
	err := os.Setenv("VIRTUAL_ENV", "/home/user/venv")
	if err != nil {
		t.Fatalf("Can't set the required enviornment.")
	}
	expected := `\[\e[48;5;22m\]  venv `
	segment := NewSegmentVirtualEnv()
	segment.Compile()
	out := segment.Render()
	if string(out) != expected {
		t.Fatalf("Rendered segment expected to be %q but got %q", expected, string(out))
	}
}
