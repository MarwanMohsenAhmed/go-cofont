package cofonts

import (
	"testing"
)

func TestParseColor(t *testing.T) {
	tests := []struct {
		input    string
		expected RGB
	}{
		{"", RGB{229, 229, 229}}, // fallback to white
		{"invalid", RGB{229, 229, 229}},
		{"red", RGB{205, 49, 49}},
		{"REDBRIGHT", RGB{241, 76, 76}}, // case insensitive
		{"#fff", RGB{255, 255, 255}},
		{"#000000", RGB{0, 0, 0}},
		{"#ff0000", RGB{255, 0, 0}},
		{"#invalid", RGB{229, 229, 229}},
	}

	for _, tt := range tests {
		actual := parseColor(tt.input)
		if actual != tt.expected {
			t.Errorf("parseColor(%q) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func TestInterpolate(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}

	tests := []struct {
		c1, c2   RGB
		t        float64
		expected RGB
	}{
		{black, white, 0.0, black},
		{black, white, 1.0, white},
		{black, white, 0.5, RGB{127, 127, 127}},
		{black, white, -0.5, black}, // bound check
		{black, white, 1.5, white},  // bound check
		{RGB{10, 20, 30}, RGB{30, 40, 50}, 0.5, RGB{20, 30, 40}},
	}

	for _, tt := range tests {
		actual := interpolate(tt.c1, tt.c2, tt.t)
		if actual != tt.expected {
			t.Errorf("interpolate(%v, %v, %v) = %v; expected %v", tt.c1, tt.c2, tt.t, actual, tt.expected)
		}
	}
}

func TestApplyBackground(t *testing.T) {
	outDefault := applyBackground("test", "transparent")
	if outDefault != "test" {
		t.Errorf("Expected 'test', got %q", outDefault)
	}

	outColor := applyBackground("test", "red")
	expected := "\x1b[48;2;205;49;49mtest\x1b[49m"
	if outColor != expected {
		t.Errorf("Expected %q, got %q", expected, outColor)
	}
}
