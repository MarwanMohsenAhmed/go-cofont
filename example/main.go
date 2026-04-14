package main

import (
	"fmt"
	"github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	// Example 1: Basic Red Text in 'block' font
	fmt.Println("--- Example 1: Basic Red Text ---")
	opts1 := cfonts.DefaultOptions()
	opts1.Font = "block"
	opts1.Colors = []string{"red"}
	res1, _ := cfonts.Render("RED", opts1)
	fmt.Print(res1)

	// Example 2: Gradient Text in 'slick' font
	fmt.Println("\n--- Example 2: Gradient Text ---")
	opts2 := cfonts.DefaultOptions()
	opts2.Font = "slick"
	opts2.Gradient = []string{"cyan", "magenta"}
	res2, _ := cfonts.Render("GRADIENT", opts2)
	fmt.Print(res2)

	// Example 3: Tiny font with Align Center
	fmt.Println("\n--- Example 3: Tiny font Centered ---")
	opts3 := cfonts.DefaultOptions()
	opts3.Font = "tiny"
	opts3.Align = "center"
	opts3.MaxLength = 40
	res3, _ := cfonts.Render("CENTERED", opts3)
	fmt.Print(res3)
}
