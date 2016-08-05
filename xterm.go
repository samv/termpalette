package termpalette

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

func XtermPaletteCode(color Color, code ...int) []byte {
	var escapeCode bytes.Buffer
	escapeCode.WriteString("\033]")
	for _, c := range code {
		escapeCode.WriteString(strconv.Itoa(c))
		escapeCode.WriteRune(';')
	}
	escapeCode.WriteString(fmt.Sprintf("rgb:%.2x/%.2x/%.2x\033\\", color[0], color[1], color[2]))
	return escapeCode.Bytes()
}

func SetThemeXterm(term io.Writer, theme Theme) {
	for i := 0; i < 16; i++ {
		term.Write(XtermPaletteCode(theme.AnsiColors[i], 4, i))
	}
	term.Write(XtermPaletteCode(theme.DefaultFg, 10))
	term.Write(XtermPaletteCode(theme.DefaultBg, 11))
	term.Write(XtermPaletteCode(theme.CursorBg, 12))
	term.Write(XtermPaletteCode(theme.SelectionBg, 17))
	term.Write(XtermPaletteCode(theme.SelectionFg, 19))
}
