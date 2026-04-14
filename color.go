package cofonts

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RGB struct {
	R, G, B uint8
	Skip    bool
}

var colorMap = map[string]RGB{
	"system":        {R: 0, G: 0, B: 0, Skip: true},
	"transparent":   {R: 0, G: 0, B: 0, Skip: true},
	"black":         {R: 0, G: 0, B: 0, Skip: false},
	"red":           {R: 205, G: 49, B: 49, Skip: false},
	"green":         {R: 13, G: 188, B: 121, Skip: false},
	"yellow":        {R: 229, G: 229, B: 16, Skip: false},
	"blue":          {R: 36, G: 114, B: 200, Skip: false},
	"magenta":       {R: 188, G: 63, B: 188, Skip: false},
	"cyan":          {R: 17, G: 168, B: 205, Skip: false},
	"white":         {R: 229, G: 229, B: 229, Skip: false},
	"gray":          {R: 102, G: 102, B: 102, Skip: false},
	"grey":          {R: 102, G: 102, B: 102, Skip: false},
	"blackbright":   {R: 102, G: 102, B: 102, Skip: false},
	"redbright":     {R: 241, G: 76, B: 76, Skip: false},
	"greenbright":   {R: 35, G: 209, B: 139, Skip: false},
	"yellowbright":  {R: 245, G: 245, B: 67, Skip: false},
	"bluebright":    {R: 59, G: 142, B: 234, Skip: false},
	"magentabright": {R: 214, G: 112, B: 214, Skip: false},
	"cyanbright":    {R: 41, G: 184, B: 219, Skip: false},
	"whitebright":   {R: 255, G: 255, B: 255, Skip: false},
}

func parseColor(s string) RGB {
	s = strings.TrimSpace(s)
	ls := strings.ToLower(s)
	if ls == "" || ls == "system" || ls == "transparent" {
		return RGB{Skip: true}
	}
	if rgb, ok := colorMap[ls]; ok {
		return rgb
	}
	if strings.HasPrefix(s, "#") {
		hex := strings.TrimPrefix(s, "#")
		if len(hex) == 3 {
			hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
		}
		if len(hex) == 6 {
			if r, err := strconv.ParseUint(hex[0:2], 16, 8); err == nil {
				if g, err := strconv.ParseUint(hex[2:4], 16, 8); err == nil {
					if b, err := strconv.ParseUint(hex[4:6], 16, 8); err == nil {
						return RGB{R: uint8(r), G: uint8(g), B: uint8(b), Skip: false}
					}
				}
			}
		}
	}
	return colorMap["white"]
}

func interpolate(c1, c2 RGB, t float64) RGB {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return RGB{
		R:    uint8(float64(c1.R) + t*(float64(c2.R)-float64(c1.R))),
		G:    uint8(float64(c1.G) + t*(float64(c2.G)-float64(c1.G))),
		B:    uint8(float64(c1.B) + t*(float64(c2.B)-float64(c1.B))),
		Skip: false,
	}
}

func applyColorTags(text string, colors []RGB) string {
	if !strings.Contains(text, "<c1>") && len(colors) > 0 {
		text = "<c1>" + text + "</c1>"
	}

	for i, c := range colors {
		tag := fmt.Sprintf("<c%d>", i+1)
		closeTag := fmt.Sprintf("</c%d>", i+1)

		if c.Skip {
			text = strings.ReplaceAll(text, tag, "")
			text = strings.ReplaceAll(text, closeTag, "")
			continue
		}
		
		ansi := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
		reset := "\x1b[39m"

		text = strings.ReplaceAll(text, tag, ansi)
		text = strings.ReplaceAll(text, closeTag, reset)
	}

	// Clean up any unmatched tags safely
	re := regexp.MustCompile(`<c\d+>|</c\d+>`)
	return re.ReplaceAllString(text, "")
}

func stripTags(text string) string {
	re := regexp.MustCompile(`<c\d+>|</c\d+>`)
	return re.ReplaceAllString(text, "")
}

func applyBackground(text string, bg string) string {
	ls := strings.ToLower(bg)
	if ls == "" || ls == "transparent" {
		return text
	}
	c := parseColor(bg)
	if c.Skip {
		return text
	}
	ansi := fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
	reset := "\x1b[49m"
	return ansi + text + reset
}
