# CoFonts-Go 🎨

```
 ██████╗ ██████╗  ██████╗ ██╗      █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔════╝██╔═══██╗██╔═══██╗██║     ██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██║     ██║   ██║██║   ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║
██║     ██║   ██║██║   ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
╚██████╗╚██████╔╝╚██████╔╝███████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
 ╚═════╝ ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
```

[![Go Reference](https://pkg.go.dev/badge/github.com/MarwanMohsenAhmed/go-cofont.svg)](https://pkg.go.dev/github.com/MarwanMohsenAhmed/go-cofont)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarwanMohsenAhmed/go-cofont)](https://goreportcard.com/report/github.com/MarwanMohsenAhmed/go-cofont)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Go Version](https://img.shields.io/github/go-mod/go-version/MarwanMohsenAhmed/go-cofont)](https://go.dev/)

**Sexy ANSI fonts for the console, written in pure Go.** A high-performance, zero-dependency port of the original [cfonts](https://github.com/dominikwilkowski/cfonts).

---

## ✨ Features

- 🚀 **Zero Dependencies** - Pure Go with `//go:embed` for all 13 fonts
- 🎨 **TrueColor Support** - 24-bit RGB gradient engine
- ⚡ **High Performance** - Benchmarked and optimized
- 🔧 **Clean API** - Simple `Render()` and `Say()` functions
- 🎬 **6 Animation Styles** - Rainbow, slide, pulse, neon, glitch, blink
- 🖥️ **Cross-Platform** - Windows, macOS, Linux support
- 📦 **Standalone Binary** - No runtime dependencies

---

## 📦 Installation

```bash
go get github.com/MarwanMohsenAhmed/go-cofont
```

---

## 🚀 Quick Start

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	opts := cofonts.DefaultOptions()
	opts.Font = "block"
	opts.Colors = []string{"cyan"}

	output, err := cofonts.Render("HELLO", opts)
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}
```

### Gradient Text

```go
opts := cofonts.DefaultOptions()
opts.Font = "slick"
opts.Gradient = []string{"#ff8800", "red"}

output, _ := cofonts.Render("GRADIENT", opts)
fmt.Print(output)
```

### Animated Logo

```go
opts := cofonts.DefaultOptions()
opts.Gradient = []string{"cyan", "magenta"}

ctx := context.Background()
cofonts.Animate(ctx, "ANIMATED", opts, cofonts.AnimationSlide, 50*time.Millisecond)
```

---

## 🎨 Available Fonts

```
┌─────────────┬──────────────┬─────────────┐
│   block     │    slick     │    tiny     │
├─────────────┼──────────────┼─────────────┤
│    grid     │   pallet     │    shade    │
├─────────────┼──────────────┼─────────────┤
│   chrome    │   simple     │ simpleblock │
├─────────────┼──────────────┼─────────────┤
│     3d      │    huge      │   console   │
└─────────────┴──────────────┴─────────────┘
```

---

## 🎬 Animation Styles

| Style | Description | Example |
|-------|-------------|---------|
| `rainbow` | Cycling rainbow colors | `cofonts.AnimationRainbow` |
| `slide` | Smooth gradient sliding | `cofonts.AnimationSlide` |
| `pulse` | Pulsing color changes | `cofonts.AnimationPulse` |
| `neon` | Neon sign effect | `cofonts.AnimationNeon` |
| `glitch` | Random glitch effects | `cofonts.AnimationGlitch` |
| `blink` | Blinking on/off | `cofonts.AnimationBlink` |

---

## ⚙️ Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `Font` | `string` | `"block"` | Font style (see fonts above) |
| `Align` | `string` | `"left"` | Text alignment: `left`, `center`, `right` |
| `Colors` | `[]string` | `["system"]` | List of colors or hex codes |
| `Gradient` | `[]string` | `nil` | Two colors for gradient (e.g., `["#f80", "red"]`) |
| `Background` | `string` | `"transparent"` | Background color |
| `LetterSpacing` | `int` | `1` | Space between characters (0-100) |
| `LineHeight` | `int` | `1` | Space between lines (1-50) |
| `MaxLength` | `int` | `0` | Max characters per line for wrapping |
| `Spaceless` | `bool` | `false` | Remove leading/trailing newlines |
| `RawMode` | `bool` | `false` | Use CRLF line endings |
| `AnimationOffset` | `float64` | `0.0` | Gradient offset (0.0 to 1.0) |

---

## 🛠️ Advanced Usage

### Multi-line Text

```go
// Use "|" to separate lines
output, _ := cofonts.Render("LINE ONE|LINE TWO", opts)
```

### Custom Colors

```go
opts.Colors = []string{"red", "#ff0055", "cyanBright"}
opts.Gradient = []string{"#00ff00", "#ff00ff"}
```

### Error Handling

```go
output, err := cofonts.Render("TEXT", opts)
if err != nil {
    // Handle invalid font or configuration
    log.Fatal(err)
}
```

---

## 📊 Benchmarks

```
BenchmarkRender-8                  2677    459552 ns/op    93252 B/op    1561 allocs/op
BenchmarkRenderGradient-8          1586    759339 ns/op   154319 B/op    2504 allocs/op
BenchmarkRenderTiny-8             10000    113049 ns/op    24151 B/op     459 allocs/op
BenchmarkRenderMultiline-8         1509    838699 ns/op   161370 B/op    3546 allocs/op
BenchmarkParseColor-8           1696030       685 ns/op       80 B/op       3 allocs/op
```

*Run benchmarks: `go test -bench=. -benchmem`*

---

## 🏗️ Project Structure

```
go-cofont/
├── animate.go          # Animation engine
├── color.go            # Color parsing & ANSI codes
├── font.go             # Font loading with embed
├── options.go          # Configuration options
├── render.go           # Core rendering logic
├── platform.go         # Cross-platform ANSI support
├── fonts/              # Embedded font files (13 fonts)
├── example/            # Usage examples
└── *_test.go           # Comprehensive tests
```

---

## 📝 Credits

This project is a Go port of the excellent [cfonts](https://github.com/dominikwilkowski/cfonts) library originally created by **Dominik Wilkowski**.

The font definitions and ASCII art layouts are property of the original author and are redistributed here under the terms of the GPL-3.0 License.

---

## 📄 License

GPL-3.0. See [LICENSE](LICENSE) for details.

---

<div align="center">

**Made with ❤️ using Go**

[Report Bug](https://github.com/MarwanMohsenAhmed/go-cofont/issues) · [Request Feature](https://github.com/MarwanMohsenAhmed/go-cofont/issues)

</div>
