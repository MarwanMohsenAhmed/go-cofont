package main

import (
	"fmt"
	"strings"

	"github.com/MarwanMohsenAhmed/go-cofont"
)

func render(text string, opts cofonts.Options) {
	fmt.Println()
	res, err := cofonts.Render(text, opts)
	if err != nil {
		fmt.Printf("Error rendering: %v\n", err)
		return
	}
	fmt.Print(res)
	fmt.Println()
}

func header(title string) {
	border := strings.Repeat("=", 80)
	fmt.Printf("\n\n%s\n", border)
	fmt.Printf("   %s\n", title)
	fmt.Printf("%s\n\n", border)
}

func main() {
	header("WELCOME TO THE FULL COFONTS SHOWCASE")
	fmt.Println("This script demonstrates all fonts, variables, and properties available.")

	// ----------------------------------------------------
	// 1. FONTS SHOWCASE
	// ----------------------------------------------------
	header("1. ALL AVAILABLE FONTS")
	fonts := []string{
		"block", "slick", "tiny", "grid", "pallet", "shade",
		"chrome", "simple", "simpleBlock", "3d", "simple3d", "huge", "console",
	}

	for i, f := range fonts {
		opts := cofonts.DefaultOptions()
		opts.Font = f
		// Alternate colors for variety
		if i%2 == 0 {
			opts.Colors = []string{"cyanBright", "blueBright"}
		} else {
			opts.Colors = []string{"yellowBright", "magentaBright"}
		}
		
		fmt.Printf("--> Font: %s\n", f)
		render(f, opts)
	}

	// ----------------------------------------------------
	// 2. ALIGNMENTS SHOWCASE
	// ----------------------------------------------------
	header("2. ALIGNMENT")
	aligns := []string{"left", "center", "right"}
	for _, a := range aligns {
		opts := cofonts.DefaultOptions()
		opts.Font = "tiny"
		opts.Align = a
		opts.Colors = []string{"greenBright"}
		render("Align "+a, opts)
	}

	// ----------------------------------------------------
	// 3. MAX LENGTH (WRAPPING)
	// ----------------------------------------------------
	header("3. MAX LENGTH (TEXT WRAP)")
	optsMax := cofonts.DefaultOptions()
	optsMax.Font = "tiny"
	optsMax.MaxLength = 20
	optsMax.Colors = []string{"redBright"}
	fmt.Println("MaxLength set to 20")
	render("This is a very long text that will wrap nicely", optsMax)

	// ----------------------------------------------------
	// 4. COLORS & BACKGROUNDS
	// ----------------------------------------------------
	header("4. COLORS AND BACKGROUNDS")
	
	optsColor := cofonts.DefaultOptions()
	optsColor.Font = "tiny"
	optsColor.Colors = []string{"red", "green", "blue"}
	fmt.Println("Multiple Colors Passed:")
	render("RGB", optsColor)

	optsHex := cofonts.DefaultOptions()
	optsHex.Font = "tiny"
	optsHex.Colors = []string{"#ff0055"}
	fmt.Println("Hex Colors (#ff0055):")
	render("Hex Color", optsHex)

	optsBg := cofonts.DefaultOptions()
	optsBg.Font = "tiny"
	optsBg.Colors = []string{"black"}
	optsBg.Background = "whiteBright"
	fmt.Println("Background Color (Black on White):")
	render("Bg Color", optsBg)

	// ----------------------------------------------------
	// 5. GRADIENTS
	// ----------------------------------------------------
	header("5. GRADIENTS")
	
	optsGradient1 := cofonts.DefaultOptions()
	optsGradient1.Font = "block"
	optsGradient1.Gradient = []string{"red", "blue"}
	fmt.Println("Simple Gradient (Red -> Blue):")
	render("Gradient", optsGradient1)

	optsGradient2 := cofonts.DefaultOptions()
	optsGradient2.Font = "block"
	optsGradient2.Gradient = []string{"red", "green", "blue"}
	fmt.Println("Multi-color Gradient (Red -> Green -> Blue):")
	render("Multi Grad", optsGradient2)

	optsGradient3 := cofonts.DefaultOptions()
	optsGradient3.Font = "block"
	optsGradient3.Gradient = []string{"yellow", "magenta"}
	optsGradient3.TransitionGradient = true
	fmt.Println("Transition Gradient (Across Multiple Lines):")
	render("Line 1\nLine 2\nLine 3", optsGradient3)
	
	optsGradient4 := cofonts.DefaultOptions()
	optsGradient4.Font = "block"
	optsGradient4.Gradient = []string{"cyan", "blue"}
	optsGradient4.IndependentGradient = true
	fmt.Println("Independent Gradient (Re-calculated per line):")
	render("Line 1\nLine 2\nLine 3", optsGradient4)

	// ----------------------------------------------------
	// 6. TYPOGRAPHY PROPS
	// ----------------------------------------------------
	header("6. TYPOGRAPHY PROPS")

	optsSpacing1 := cofonts.DefaultOptions()
	optsSpacing1.Font = "tiny"
	optsSpacing1.LetterSpacing = 3
	optsSpacing1.Colors = []string{"cyanBright"}
	fmt.Println("LetterSpacing: 3 (Distance between letters)")
	render("SPACED", optsSpacing1)

	optsLineHeight := cofonts.DefaultOptions()
	optsLineHeight.Font = "tiny"
	optsLineHeight.LineHeight = 3
	optsLineHeight.Colors = []string{"magentaBright"}
	fmt.Println("LineHeight: 3 (Distance between lines)")
	render("Line 1\nLine 2", optsLineHeight)

	optsSpaceless := cofonts.DefaultOptions()
	optsSpaceless.Font = "block"
	optsSpaceless.Spaceless = true
	optsSpaceless.Colors = []string{"greenBright"}
	fmt.Println("Spaceless: True (No extra padding between letters)")
	render("SPACE", optsSpaceless)

	header("SHOWCASE COMPLETE")
}
