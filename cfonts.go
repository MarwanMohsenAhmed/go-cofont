package cfonts

import (
	"fmt"
)

// Render computes the beautiful ANSI string output based on default options.
func Render(text string, opts Options) (string, error) {
	if text == "" {
		return "", nil
	}

	fontSchema, err := GetFont(opts.Font)
	if err != nil {
		return "", err
	}

	return renderCore(text, opts, fontSchema), nil
}

// Say prints the rendered cfonts sequence to standard output.
func Say(text string, opts Options) error {
	out, err := Render(text, opts)
	if err != nil {
		return err
	}
	fmt.Print(out)
	return nil
}
