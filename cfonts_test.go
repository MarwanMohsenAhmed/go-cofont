package cfonts

import (
	"strings"
	"testing"
)

func TestRenderBasic(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.Colors = []string{"red"}

	out, err := Render("A", opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// For tiny font 'A', the char string looks like:
	// "▄▀█"
	// "█▀█"
	
	// Check if red color is applied
	if !strings.Contains(out, "\x1b[38;2;205;49;49m") {
		t.Errorf("Expected output to contain red color escape sequences, got: %s", out)
	}
	
	// Check if char is rendered
	if !strings.Contains(out, "▄▀█") || !strings.Contains(out, "█▀█") {
		t.Errorf("Expected output to contain 'A' shape, got: %s", out)
	}
}

func TestRenderGradient(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.Gradient = []string{"red", "blue"}

	out, err := Render("AB", opts)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Should contain the start red and end blue gradients approximately
	if !strings.Contains(out, "\x1b[38;2;") {
		t.Errorf("Expected output to contain gradient color escape sequences")
	}
}

func TestGetFont(t *testing.T) {
	_, err := GetFont("chrome")
	if err != nil {
		t.Fatalf("Expected to find 'chrome' font, got error: %v", err)
	}
	
	_, err = GetFont("nonexistent")
	if err == nil {
		t.Fatalf("Expected error for 'nonexistent' font, got nil")
	}
}
