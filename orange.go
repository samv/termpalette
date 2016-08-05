package termpalette

var Orange = Theme{
	DefaultBg:   Color{238, 238, 209}, // LightYellow2
	DefaultFg:   Color{18, 18, 18},    // gray7
	BoldFg:      Color{255, 69, 0},    // OrangeRed
	CursorBg:    Color{83, 134, 139},  // CadetBlue4
	CursorFg:    Color{238, 223, 204}, // AntiqueWhite2
	SelectionFg: Color{139, 54, 38},   // tomato4
	SelectionBg: Color{205, 183, 181}, // MistyRose3
	AnsiColors: [16]Color{
		Color{31, 31, 31},    // grey12
		Color{139, 26, 26},   // firebrick4
		Color{69, 139, 0},    // chartreuse4
		Color{139, 101, 8},   // DarkGoldenrod4
		Color{16, 78, 139},   // DodgerBlue4
		Color{104, 34, 139},  // DarkOrchid4
		Color{0, 134, 139},   // turquoise4
		Color{205, 205, 180}, // LightYellow3
		Color{23, 23, 23},    // grey9
		Color{205, 0, 0},     // red3
		Color{0, 139, 0},     // green4
		Color{205, 149, 12},  // DarkGoldenrod3
		Color{79, 148, 205},  // SteelBlue3
		Color{255, 105, 180}, // HotPink
		Color{0, 206, 209},   // DarkTurquoise
		Color{255, 140, 0},   // DarkOrange
	},
}
