package cofonts

import (
	"testing"
)

func BenchmarkRender(b *testing.B) {
	opts := DefaultOptions()
	opts.Font = "block"
	text := "HELLO"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Render(text, opts)
	}
}

func BenchmarkRenderGradient(b *testing.B) {
	opts := DefaultOptions()
	opts.Font = "slick"
	opts.Gradient = []string{"cyan", "magenta"}
	text := "GRADIENT"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Render(text, opts)
	}
}

func BenchmarkRenderTiny(b *testing.B) {
	opts := DefaultOptions()
	opts.Font = "tiny"
	text := "ABC"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Render(text, opts)
	}
}

func BenchmarkRenderMultiline(b *testing.B) {
	opts := DefaultOptions()
	opts.Font = "block"
	opts.Spaceless = true
	text := "LINE ONE|LINE TWO"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Render(text, opts)
	}
}

func BenchmarkRenderWithColors(b *testing.B) {
	opts := DefaultOptions()
	opts.Font = "chrome"
	opts.Colors = []string{"red", "green", "blue"}
	text := "COLORS"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Render(text, opts)
	}
}

func BenchmarkParseColor(b *testing.B) {
	colors := []string{"red", "#ff0055", "cyanBright", "invalid", "#fff"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range colors {
			_ = parseColor(c)
		}
	}
}
