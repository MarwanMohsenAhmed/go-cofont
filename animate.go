package cofonts

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type AnimationStyle string

const (
	AnimationRainbow AnimationStyle = "rainbow"
	AnimationSlide   AnimationStyle = "slide"
	AnimationPulse   AnimationStyle = "pulse"
	AnimationNeon    AnimationStyle = "neon"
	AnimationGlitch  AnimationStyle = "glitch"
	AnimationBlink   AnimationStyle = "blink"
)

// ValidAnimationStyles returns all supported animation styles.
func ValidAnimationStyles() []AnimationStyle {
	return []AnimationStyle{
		AnimationRainbow,
		AnimationSlide,
		AnimationPulse,
		AnimationNeon,
		AnimationGlitch,
		AnimationBlink,
	}
}

// IsValid checks if the animation style is supported.
func (s AnimationStyle) IsValid() bool {
	switch s {
	case AnimationRainbow, AnimationSlide, AnimationPulse,
		AnimationNeon, AnimationGlitch, AnimationBlink:
		return true
	}
	return false
}

// animationCycleCount controls the gradient offset cycle for slide animation.
const animationCycleCount = 100

var (
	rainbowColors = []string{"red", "yellow", "green", "cyan", "blue", "magenta"}
	neonColors    = []string{"cyanBright", "magentaBright", "yellowBright", "redBright"}
)

// Animate a logo in the terminal until context is cancelled or an error occurs.
func Animate(ctx context.Context, text string, opts Options, style AnimationStyle, speed time.Duration) error {
	// Validate animation style
	if !style.IsValid() {
		return fmt.Errorf("unsupported animation style: %q (valid: %v)", style, ValidAnimationStyles())
	}

	// Validate speed to prevent 100% CPU usage
	if speed <= 0 {
		speed = 100 * time.Millisecond
	}

	// Hide cursor
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h") // Restore cursor on exit

	frame := 0
	maxLineCount := 0

	// Internal function to render and print a single frame
	printFrame := func() error {
		// Deep-copy options to prevent race conditions on slices
		currentOpts := opts
		currentOpts.Colors = make([]string, len(opts.Colors))
		copy(currentOpts.Colors, opts.Colors)
		if opts.Gradient != nil {
			currentOpts.Gradient = make([]string, len(opts.Gradient))
			copy(currentOpts.Gradient, opts.Gradient)
		}

		switch style {
		case AnimationRainbow:
			cIdx := frame % len(rainbowColors)
			currentOpts.Colors = []string{rainbowColors[cIdx]}
			if len(currentOpts.Gradient) >= 2 {
				c2Idx := (frame + 1) % len(rainbowColors)
				currentOpts.Gradient = []string{rainbowColors[cIdx], rainbowColors[c2Idx]}
			}
		case AnimationSlide:
			currentOpts.AnimationOffset = float64(frame%animationCycleCount) / float64(animationCycleCount)
			if len(currentOpts.Gradient) < 2 {
				currentOpts.Gradient = []string{"red", "blue"}
			}
		case AnimationPulse:
			cIdx := (frame / 4) % len(rainbowColors)
			currentOpts.Colors = []string{rainbowColors[cIdx]}
		case AnimationNeon:
			cIdx := (frame / 2) % len(neonColors)
			currentOpts.Colors = []string{neonColors[cIdx]}
		case AnimationBlink:
			if frame%2 == 0 {
				currentOpts.Colors = []string{"system"}
				currentOpts.Background = "transparent"
				currentOpts.Gradient = nil
			}
		case AnimationGlitch:
			cIdx := rand.Intn(len(rainbowColors))
			currentOpts.Colors = []string{rainbowColors[cIdx]}
			if rand.Float32() < 0.2 {
				currentOpts.AnimationOffset = rand.Float64()
			}
		}

		// Render the frame
		out, err := Render(text, currentOpts)
		if err != nil {
			return err
		}

		// Clear previous frame by tracking max line count
		// This prevents ghosting when line counts vary between frames
		clearLines := maxLineCount
		if lineCount := strings.Count(out, "\n"); lineCount > clearLines {
			clearLines = lineCount
		}

		if clearLines > 0 {
			// Move cursor up and clear to end of screen
			fmt.Printf("\033[%dF", clearLines)
			fmt.Print("\033[J")
		} else {
			fmt.Print("\r")
		}

		// Print new frame
		fmt.Print(out)

		// Update max line count for next iteration
		if lineCount := strings.Count(out, "\n"); lineCount > maxLineCount {
			maxLineCount = lineCount
		}

		frame++
		return nil
	}

	// 1. Initial frame immediately
	if err := printFrame(); err != nil {
		return err
	}

	// 2. Loop
	ticker := time.NewTicker(speed)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println()
			return nil
		case <-ticker.C:
			if err := printFrame(); err != nil {
				return err
			}
		}
	}
}
