package cofonts

type Options struct {
	Font                string
	Align               string
	Colors              []string
	Gradient            []string
	IndependentGradient bool
	TransitionGradient  bool
	Background          string
	LetterSpacing       int
	LineHeight          int
	Spaceless           bool
	MaxLength           int
	RawMode             bool
	Env                 string
	AnimationOffset     float64
}

// DefaultOptions returns a set of default options per the js/rust library.
func DefaultOptions() Options {
	return Options{
		Font:                "block",
		Align:               "left",
		Colors:              []string{"system"},
		Gradient:            nil,
		IndependentGradient: false,
		TransitionGradient:  false,
		Background:          "transparent",
		LetterSpacing:       1,
		LineHeight:          1,
		Spaceless:           false,
		MaxLength:           0,
		RawMode:             false,
		Env:                 "cli",
	}
}
