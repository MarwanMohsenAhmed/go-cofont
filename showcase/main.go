package main

import (
	"fmt"
	"github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	fonts := []string{
		"block", "slick", "tiny", "grid", "pallet", "shade",
		"chrome", "simple", "simpleBlock", "3d", "simple3d", "huge", "console",
	}

	fmt.Println("========================================")
	fmt.Println("    WELCOME TO THE COFONTS SHOWCASE     ")
	fmt.Println("========================================")

	for i, f := range fonts {
		opts := cofonts.DefaultOptions()
		opts.Font = f
		
		// Cycle some colors and gradients for fun
		if i%3 == 0 {
			opts.Colors = []string{"greenBright"}
		} else if i%3 == 1 {
			opts.Gradient = []string{"cyan", "blue"}
		} else {
			opts.Gradient = []string{"red", "#ffaa00"}
		}

		res, err := cofonts.Render("F: "+f, opts)
		if err != nil {
			fmt.Printf("\nError rendering %s: %v\n", f, err)
			continue
		}
		
		// Print it!
		fmt.Printf("\n--- Font: %s ---\n", f)
		fmt.Print(res)
	}
	
	fmt.Println("\n========================================")
	fmt.Println("       SHOWCASE COMPLETE                ")
	fmt.Println("========================================")
}
