package cofonts

import (
	"os"
	"runtime"
	"strings"
)

// supportsANSI checks if the current terminal supports ANSI escape sequences.
// On modern systems (Windows 10+, macOS, Linux), this is typically true.
func supportsANSI() bool {
	// Windows detection
	if runtime.GOOS == "windows" {
		// Check for WT_SESSION (Windows Terminal) or TERM (MSYS/Cygwin/Git Bash)
		if os.Getenv("WT_SESSION") != "" {
			return true
		}
		if strings.Contains(os.Getenv("TERM"), "xterm") {
			return true
		}
		// Windows 10+ build 14393+ supports ANSI via Virtual Terminal
		// We assume modern Go installations run on Windows 10+
		return true
	}

	// Unix-like systems generally support ANSI
	return true
}

// escapeSequence returns the appropriate escape sequence prefix.
func escapeSequence() string {
	if supportsANSI() {
		return "\x1b"
	}
	// Fallback for legacy systems: return empty string (no colors)
	return ""
}
