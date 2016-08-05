package termpalette

import (
	"io"
	"os"
	"strings"
)

type Color [3]uint8

type Theme struct {
	AnsiColors [16]Color
	BoldFg     Color
	//Underline Color
	DefaultBg   Color
	DefaultFg   Color
	CursorFg    Color
	CursorBg    Color
	SelectionFg Color
	SelectionBg Color
}

// SetTheme maybe writes control codes to set the color theme on the
// terminal, using the value of TERM and/or a couple of variables sent
// by iTerm2 to know whether to send iTerm2 escape codes or more
// standard xterm escape codes.  It will only work with iTerm if the
func SetTheme(term io.Writer, theme Theme) {
	var isXterm, isIterm bool

	// TERM suffers from the same problem as browser User-Agent strings, except they
	// just all say they're xterm or xterm-256color, so use other hints.

	// Terminal detection would be better but I'm already far too far down this rabbit hole :)
	termName := os.Getenv("TERM")
	iterm_profile := os.Getenv("ITERM_PROFILE")
	xterm_version := os.Getenv("XTERM_VERSION")

	if strings.Contains(termName, "xterm") || strings.Contains(termName, "256color") {
		isXterm = true
	}
	if iterm_profile != "" {
		isIterm = true
		isXterm = false
	}
	if xterm_version != "" {
		isXterm = true
	}
	if isXterm {
		SetThemeXterm(term, theme)
	} else if isIterm {
		SetThemeIterm(term, theme)
	}
}

// SetDarkTheme sets a theme with dark background and green text.
func SetDarkTheme(term io.Writer) {
	SetTheme(term, HipsterGreen)
}

// SetLightTheme sets a theme with an off-white background and dark text
func SetLightTheme(term io.Writer) {
	SetTheme(term, Orange)
}
