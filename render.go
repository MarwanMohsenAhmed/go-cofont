package cofonts

import (
	"math"
	"strings"
	"unicode/utf8"
)

func renderCore(text string, opts Options, font *FontSchema) string {
	// First convert explicit 'new lines' pipe symbol
	text = strings.ReplaceAll(text, "|", "\n")
	parts := strings.Split(text, "\n")

	var outputLines []string

	// Optional gradient parsing
	useGradient := len(opts.Gradient) >= 2
	var gStart, gEnd RGB
	if useGradient {
		gStart = parseColor(opts.Gradient[0])
		gEnd = parseColor(opts.Gradient[1])
	}
	
	// Default base colors if no gradient
	baseColors := make([]RGB, font.Colors)
	for i := 0; i < font.Colors; i++ {
		if i < len(opts.Colors) {
			baseColors[i] = parseColor(opts.Colors[i])
		} else if len(opts.Colors) > 0 {
			// Fallback to first color if not enough colors provided
			baseColors[i] = parseColor(opts.Colors[0])
		} else {
			baseColors[i] = parseColor("system")
		}
	}

	for lineIdx, lineStr := range parts {
		if lineStr == "" {
			continue
		}

		charArrays := make([][]string, 0, len(lineStr))
		for _, r := range lineStr {
			charStr := strings.ToUpper(string(r))
			c, ok := font.Chars[charStr]
			if !ok {
				c, ok = font.Chars[" "]
				if !ok {
					// Hard fallback
					c = make([]string, font.Lines)
					for i := range c {
						c[i] = " "
					}
				}
			}
			charArrays = append(charArrays, c)
		}

		lineRows := make([]string, font.Lines)
		
		totalWidth := 0
		for _, c := range charArrays {
			// Length of the first line of the char, striped of tags
			totalWidth += utf8.RuneCountInString(stripTags(c[0]))
			totalWidth += font.LetterspaceSize
		}
		if len(charArrays) > 0 {
			totalWidth -= font.LetterspaceSize
		}

		for r := 0; r < font.Lines; r++ {
			rowBuilder := strings.Builder{}
			currentX := 0
			
			for cIdx, charLines := range charArrays {
				charStr := charLines[r]
				
				// Gradient color calculation if needed
				var renderColors []RGB
				if useGradient {
					// We calculate color mix based on exact X position 
					progress := 0.0
					if totalWidth > 1 {
						progress = float64(currentX) / float64(totalWidth-1)
					}
					
					// Simple horizontal gradient
					singleMixedColor := interpolate(gStart, gEnd, progress)
					
					// Replicate mixed color for all `<cN>` tags inside this character
					renderColors = make([]RGB, font.Colors)
					for ci := range renderColors {
						renderColors[ci] = singleMixedColor
					}
					
				} else {
					renderColors = baseColors
				}

				coloredChar := applyColorTags(charStr, renderColors)
				rowBuilder.WriteString(coloredChar)
				
				currentX += utf8.RuneCountInString(stripTags(charStr))

				// Add letter spacing except for last char
				if cIdx < len(charArrays)-1 {
					lsStr := ""
					if r < len(font.Letterspace) {
						lsStr = font.Letterspace[r]
					} else {
						lsStr = strings.Repeat(" ", font.LetterspaceSize)
					}
					// Also apply color to letter spacing
					coloredLs := applyColorTags(lsStr, renderColors)
					rowBuilder.WriteString(coloredLs)
					currentX += font.LetterspaceSize
				}
			}
			
			// Apply alignment
			finalStr := rowBuilder.String()
			strippedFinal := stripTags(finalStr)
			actualLen := utf8.RuneCountInString(strippedFinal)
			
			padLeft := 0
			if opts.MaxLength > 0 && actualLen < opts.MaxLength {
				if opts.Align == "center" {
					padLeft = int(math.Max(0, float64(opts.MaxLength-actualLen)/2))
				} else if opts.Align == "right" {
					padLeft = opts.MaxLength - actualLen
				}
			}

			if padLeft > 0 {
				finalStr = strings.Repeat(" ", padLeft) + finalStr
			}

			lineRows[r] = applyBackground(finalStr, opts.Background)
		}

		outputLines = append(outputLines, lineRows...)

		// apply line height (empty lines between logical lines)
		if lineIdx < len(parts)-1 && opts.LineHeight > 1 {
			for i := 1; i < opts.LineHeight; i++ {
				outputLines = append(outputLines, "")
			}
		}
	}

	joinStr := "\n"
	if opts.RawMode {
		joinStr = "\r\n"
	}

	result := strings.Join(outputLines, joinStr)

	if !opts.Spaceless {
		result = joinStr + joinStr + result + joinStr + joinStr
	}

	return result
}
