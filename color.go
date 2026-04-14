package cofonts

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RGB struct {
	R, G, B uint8
}

var colorMap = map[string]RGB{
	"system":        {0, 0, 0}, // Fallback marker
	"transparent":   {0, 0, 0},
	"black":         {0, 0, 0},
	"red":           {205, 49, 49},
	"green":         {13, 188, 121},
	"yellow":        {229, 229, 16},
	"blue":          {36, 114, 200},
	"magenta":       {188, 63, 188},
	"cyan":          {17, 168, 205},
	"white":         {229, 229, 229},
	"gray":          {102, 102, 102},
	"grey":          {102, 102, 102},
	"blackBright":   {102, 102, 102},
	"redBright":     {241, 76, 76},
	"greenBright":   {35, 209, 139},
	"yellowBright":  {245, 245, 67},
	"blueBright":    {59, 142, 234},
	"magentaBright": {214, 112, 214},
	"cyanBright":    {41, 184, 219},
	"whiteBright":   {255, 255, 255},
}

func parseColor(s string) RGB {
	s = strings.TrimSpace(s)
	if s == "" {
		return colorMap["white"]
	}
	if rgb, ok := colorMap[strings.ToLower(s)]; ok {
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
						return RGB{uint8(r), uint8(g), uint8(b)}
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
		R: uint8(float64(c1.R) + t*(float64(c2.R)-float64(c1.R))),
		G: uint8(float64(c1.G) + t*(float64(c2.G)-float64(c1.G))),
		B: uint8(float64(c1.B) + t*(float64(c2.B)-float64(c1.B))),
	}
}

func applyColorTags(text string, colors []RGB) string {
	if !strings.Contains(text, "<c1>") && len(colors) > 0 {
		text = "<c1>" + text + "</c1>"
	}

	for i, c := range colors {
		tag := fmt.Sprintf("<c%d>", i+1)
		closeTag := fmt.Sprintf("</c%d>", i+1)
		
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
	ansi := fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
	reset := "\x1b[49m"
	return ansi + text + reset
}
