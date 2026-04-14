package main

import (
	"fmt"
	cofonts "github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	opts := cofonts.DefaultOptions()
	opts.Font = "block"
	opts.Colors = []string{"system"} // avoid ansi codes
	res, _ := cofonts.Render("Test", opts)
	fmt.Print(res)
}
