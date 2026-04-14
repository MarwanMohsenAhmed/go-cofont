package cofonts

import (
	"strings"
	"testing"
)

func TestAlignmentOptions(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.MaxLength = 50
	opts.Spaceless = true

	// With "tiny" font, "A" produces something approximately 3 columns wide.
	outLeft, _ := Render("A", opts)
	opts.Align = "right"
	outRight, _ := Render("A", opts)
	opts.Align = "center"
	_, _ = Render("A", opts) // Ensure it doesn't crash on center

	// Since they all contain padding differently:
	// A strictly left aligned line string starts with ansi tags or the block directly.
	// A strictly right aligned line string starts with many spaces.
	// A strictly centered line string starts with roughly half the spaces of right aligned.
	
	// We check standard byte prefixes. The literal space " " prefix indicates alignment padding.
	leftLines := strings.Split(outLeft, "\x1b[39m") // Naive split simply to get rows roughly
	rightLines := strings.Split(outRight, "\x1b")
	
	if len(leftLines) > 0 && len(rightLines) > 0 {
		leftRow1 := leftLines[0]
		rightRow1 := rightLines[0]
		
		// The right aligned output MUST contain significantly more literal space padding than the left.
		leftSpaceCount := strings.Count(leftRow1, " ")
		rightSpaceCount := strings.Count(rightRow1, " ")
		
		if rightSpaceCount <= leftSpaceCount {
			t.Errorf("Expected right-aligned to contain more padding spaces than left-aligned. Right: %d, Left: %d", rightSpaceCount, leftSpaceCount)
		}
	}
}

func TestLetterSpacingAndLineHeight(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.Spaceless = true

	outStandard, _ := Render("A", opts)
	
	opts.LineHeight = 3
	outHeight, _ := Render("A|B", opts) // Pipe adds a new logical line
	
	// A standard font produces 2 lines for A, and 2 for B = 4 lines.
	// LineHeight = 3 adds (3-1=2) blank rows between logical lines.
	// Total expected lines = 4 + 2 = 6 lines.
	lines := strings.Split(strings.TrimSpace(outHeight), "\n")
	if len(lines) != 6 {
		t.Errorf("Expected 6 lines total for LineHeight=3 between A and B, got %d", len(lines))
	}
	
	_ = outStandard
}

func TestSpaceless(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	
	outPad, _ := Render("A", opts)
	
	opts.Spaceless = true
	outNoPad, _ := Render("A", opts)
	
	if len(outPad) <= len(outNoPad) {
		t.Errorf("Expected Spaceless=false to produce a larger string due to newlines")
	}
	
	if strings.HasPrefix(outNoPad, "\n\n") || strings.HasSuffix(outNoPad, "\n\n") {
		t.Errorf("Expected Spaceless=true output to omit leading/trailing newline blocks")
	}
}

func TestRawMode(t *testing.T) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	opts.RawMode = true
	opts.Spaceless = true

	out, _ := Render("A", opts)
	if !strings.Contains(out, "\r\n") {
		t.Errorf("Expected RawMode=true output to contain CRLF (\\r\\n), got strings without it.")
	}
}
