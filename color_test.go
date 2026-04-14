package cofonts

import (
	"testing"
)

func TestParseColor(t *testing.T) {
	tests := []struct {
		input         string
		expectedRGB   RGB
		expectWarning bool
	}{
		{"", RGB{Skip: true}, false},
		{"system", RGB{Skip: true}, false},
		{"transparent", RGB{Skip: true}, false},
		{"invalid", RGB{R: 229, G: 229, B: 229, Skip: false}, true}, // unknown color
		{"red", RGB{R: 205, G: 49, B: 49, Skip: false}, false},
		{"REDBRIGHT", RGB{R: 241, G: 76, B: 76, Skip: false}, false}, // case insensitive
		{"#fff", RGB{R: 255, G: 255, B: 255, Skip: false}, false},
		{"#000000", RGB{R: 0, G: 0, B: 0, Skip: false}, false},
		{"#ff0000", RGB{R: 255, G: 0, B: 0, Skip: false}, false},
		{"#GGGGGG", RGB{R: 229, G: 229, B: 229, Skip: false}, true}, // invalid hex
		{"#invalid", RGB{R: 229, G: 229, B: 229, Skip: false}, true}, // invalid format
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseColor(tt.input)
			if result.RGB != tt.expectedRGB {
				t.Errorf("parseColor(%q) RGB = %v; expected %v", tt.input, result.RGB, tt.expectedRGB)
			}
			hasWarning := result.Warning != ""
			if hasWarning != tt.expectWarning {
				t.Errorf("parseColor(%q) warning = %q; expected warning: %v", tt.input, result.Warning, tt.expectWarning)
			}
		})
	}
}

func TestInterpolate(t *testing.T) {
	black := RGB{R: 0, G: 0, B: 0, Skip: false}
	white := RGB{R: 255, G: 255, B: 255, Skip: false}

	tests := []struct {
		c1, c2   RGB
		t        float64
		expected RGB
	}{
		{black, white, 0.0, black},
		{black, white, 1.0, white},
		{black, white, 0.5, RGB{R: 127, G: 127, B: 127, Skip: false}},
		{black, white, -0.5, black}, // bound check
		{black, white, 1.5, white},  // bound check
		{RGB{R: 10, G: 20, B: 30, Skip: false}, RGB{R: 30, G: 40, B: 50, Skip: false}, 0.5, RGB{R: 20, G: 30, B: 40, Skip: false}},
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

func TestParseColorWarnings(t *testing.T) {
	// Test that invalid colors produce warnings
	result := parseColor("#ZZZZZZ")
	if result.Warning == "" {
		t.Error("Expected warning for invalid hex color #ZZZZZZ")
	}

	result2 := parseColor("blurple")
	if result2.Warning == "" {
		t.Error("Expected warning for unknown color 'blurple'")
	}
}
