package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestMain(t *testing.T) {
	out, in := makeUI()
	
	if out.Text != "Hello" {
		t.Errorf("Expected label text to be 'Hello', got '%s'", out.Text)
	}

	test.Type(in, "world")
	if out.Text != "Hello world" {
		t.Errorf("Expected label text to be 'Hello world', got '%s'", out.Text)
	}
}