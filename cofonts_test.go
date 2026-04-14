package cofonts

import (
	"strings"
	"testing"
)

var allFonts = []string{
	"3d", "block", "chrome", "console", "grid", "huge",
	"pallet", "shade", "simple", "simple3d", "simpleBlock", "slick", "tiny",
}

func TestAllFontsRender(t *testing.T) {
	opts := DefaultOptions()
	text := "A1"

	for _, fontName := range allFonts {
		t.Run(fontName, func(t *testing.T) {
			opts.Font = fontName
			out, err := Render(text, opts)
			if err != nil {
				t.Fatalf("Render failed for %s: %v", fontName, err)
			}
			if out == "" {
				t.Fatalf("Render returned empty string for %s", fontName)
			}

			// Validate we have line breaks (output should be multiline)
			if !strings.Contains(out, "\n") {
				t.Fatalf("Expected multiline output for font %s, got: %q", fontName, out)
			}

			// Get the font schema to compare line counts
			schema, _ := GetFont(fontName)
			
			// We expect the text to be rendered. Since 'A' is universally defined,
			// its internal tags stripped representation will be present, but ansi tags exist.
			
			// By default Spaceless is false, so it adds two lines before and after.
			// Plus the actual character lines.
			parts := strings.Split(out, "\n")
			expectedMinLines := schema.Lines + 4 // 2 top + 2 bottom padding
			if len(parts) < expectedMinLines {
				t.Errorf("Font %s output too short. Expected >= %d lines, got %d", fontName, expectedMinLines, len(parts))
			}
		})
	}
}

func TestInvalidFont(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "imaginaryFontXYZ"

	_, err := Render("test", opts)
	if err == nil {
		t.Fatal("Expected error when providing an invalid font name")
	}
	if !strings.Contains(err.Error(), "font not found") {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestUnsupportedChars(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "block"
	// ~ is usually an unsupported character in cfonts
	out, err := Render("~", opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Missing char falls back to space map -> prints empty blocks
	if out == "" {
		t.Error("Render returned empty strings for unsupported char")
	}
}

func TestEmptyInput(t *testing.T) {
	opts := DefaultOptions()
	out, err := Render("", opts)
	if err != nil {
		t.Fatalf("Expected nil error for empty string, got %v", err)
	}
	if out != "" {
		t.Errorf("Expected empty output for empty input, got %q", out)
	}
}

func TestPipeNewlines(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.Spaceless = true // remove the 4 extra padding lines
	
	out, err := Render("A|B", opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	// A is 2 lines, B is 2 lines. 
	// Total expected logical lines printed: 4 lines. We split by \n.
	// Since tiny has 2 lines per character.
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 4 {
		t.Errorf("Expected 4 lines for stacked A and B, got %d lines: \n%s", len(lines), out)
	}
}
