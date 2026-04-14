# CoFonts-Go

[![Go Reference](https://pkg.go.dev/badge/github.com/MarwanMohsenAhmed/go-cofont.svg)](https://pkg.go.dev/github.com/MarwanMohsenAhmed/go-cofont)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarwanMohsenAhmed/go-cofont)](https://goreportcard.com/report/github.com/MarwanMohsenAhmed/go-cofont)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Sexy ANSI fonts for the console, written in pure Go. This is a high-performance, zero-dependency port of the original [cfonts](https://github.com/dominikwilkowski/cfonts).

## Features
- **Zero Dependencies**: Uses `//go:embed` to ship all 13 fonts inside the binary.
- **Pure Go**: No CGO or external color libraries required.
- **TrueColor Support**: Native 24-bit RGB gradient engine.
- **Clean Room Port**: Avoids original JS/Rust licensing entanglements while maintaining visual parity.

## Installation

```bash
go get github.com/MarwanMohsenAhmed/go-cofont
```

## Usage

### Simple Render
```go
package main

import (
	"fmt"
	"github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	opts := cofonts.DefaultOptions()
	opts.Font = "block"
	opts.Gradient = []string{"#ff8800", "red"}

	output, err := cofonts.Render("GO BLAZE", opts)
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}
```

### Real-time Animation
```go
package main

import (
	"context"
	"time"
	"github.com/MarwanMohsenAhmed/go-cofont"
)

func main() {
	opts := cofonts.DefaultOptions()
	opts.Gradient = []string{"cyan", "magenta"}
	
	// Animate until context is cancelled
	ctx := context.Background()
	cofonts.Animate(ctx, "ANIMATED", opts, cofonts.AnimationSlide, 50*time.Millisecond)
}
```

## Options

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `Font` | `string` | `"block"` | Font style: `block`, `slick`, `tiny`, `grid`, `pallet`, `shade`, `chrome`, etc. |
| `Align` | `string` | `"left"` | Alignment: `left`, `center`, `right`. |
| `Colors` | `[]string`| `["system"]`| List of colors for the font. |
| `Gradient` | `[]string`| `nil` | Two colors for horizontal gradient (e.g. `["#f80", "red"]`). |
| `AnimationOffset` | `float64` | `0.0` | Internal offset (0.0 to 1.0) for shifting gradients manually. |
| `LetterSpacing`| `int` | `1` | Space between characters. |
| `LineHeight` | `int` | `1` | Space between lines. |
| `MaxLength` | `int` | `0` | Max characters per line for wrapping. |
| `TransitionGradient` | `bool` | `false`| Enable smooth color transitions. |
| `IndependentGradient`| `bool` | `false`| Apply gradient per line independently. |

## Credits
This project is a Go port of the excellent [cfonts](https://github.com/dominikwilkowski/cfonts) 
library originally created by **Dominik Wilkowski**. 

The font definitions and ASCII art layouts used in this library are property 
of the original author and are redistributed here under the terms of the GPL-3.0 License.

## License
GPL-3.0. See [LICENSE](LICENSE) for details.
