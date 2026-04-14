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

var (
	rainbowColors = []string{"red", "yellow", "green", "cyan", "blue", "magenta"}
	neonColors    = []string{"cyanBright", "magentaBright", "yellowBright", "redBright"}
)

// Animate a logo in the terminal until context is cancelled or an error occurs.
func Animate(ctx context.Context, text string, opts Options, style AnimationStyle, speed time.Duration) error {
	// Hide cursor
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h") // Restore cursor on exit

	frame := 0
	lineCount := 0
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Internal function to render and print a single frame
	printFrame := func() error {
		currentOpts := opts // Copy original options

		switch style {
		case AnimationRainbow:
			cIdx := frame % len(rainbowColors)
			currentOpts.Colors = []string{rainbowColors[cIdx]}
			if len(currentOpts.Gradient) >= 2 {
				c2Idx := (frame + 1) % len(rainbowColors)
				currentOpts.Gradient = []string{rainbowColors[cIdx], rainbowColors[c2Idx]}
			}
		case AnimationSlide:
			currentOpts.AnimationOffset = float64(frame%100) / 100.0
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
			cIdx := rng.Intn(len(rainbowColors))
			currentOpts.Colors = []string{rainbowColors[cIdx]}
			if rng.Float32() < 0.2 {
				currentOpts.AnimationOffset = rng.Float64()
			}
		}

		// Render the frame
		out, err := Render(text, currentOpts)
		if err != nil {
			return err
		}

		// Clear previous frame by moving up to its first line
		if lineCount > 0 {
			// \033[F moves to start of line 1 line up.
			// \033[NF moves to start of line N lines up.
			fmt.Printf("\033[%dF", lineCount)
			// Optional: Clear to end of screen to avoid ghosting
			fmt.Print("\033[J")
		} else {
			// First frame: just ensure we are at the start of the line
			fmt.Print("\r")
		}

		// Print new frame
		fmt.Print(out)
		
		// Update line count for next iteration
		// Each \n in the output string represents a transition to a new line.
		// If the output was "\n\nTEXT\n\n", strings.Count is 4. 
		// That means the terminal cursor is 4 lines below where it started.
		lineCount = strings.Count(out, "\n")
		
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
