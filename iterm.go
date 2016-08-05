package termpalette

import (
	"fmt"
	"io"
)

func ItermPaletteCode(code string, color Color) []byte {
	return []byte(fmt.Sprintf("\033]P%s%.2x%.2x%.2x\033\\", code, color[0], color[1], color[2]))
}

func SetThemeIterm(term io.Writer, theme Theme) {
	for i := 0; i < 16; i++ {
		term.Write(ItermPaletteCode(fmt.Sprintf("%x", i), theme.AnsiColors[i]))
	}
	term.Write(ItermPaletteCode("g", theme.DefaultFg))
	term.Write(ItermPaletteCode("h", theme.DefaultBg))
	term.Write(ItermPaletteCode("i", theme.BoldFg))
	term.Write(ItermPaletteCode("j", theme.SelectionBg))
	term.Write(ItermPaletteCode("k", theme.SelectionFg))
	term.Write(ItermPaletteCode("l", theme.CursorBg))
	term.Write(ItermPaletteCode("m", theme.CursorFg))
}
