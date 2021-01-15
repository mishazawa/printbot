package printer

import (
	"image/color"

	"github.com/esimov/colorquant"
)

/*
	This filters was borrowed from following repository
	https://github.com/esimov/colorquant
*/

var floydSteinberg colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 7.0 / 48.0, 5.0 / 48.0},
		{3.0 / 48.0, 5.0 / 48.0, 7.0 / 48.0, 5.0 / 48.0, 3.0 / 48.0},
		{1.0 / 48.0, 3.0 / 48.0, 5.0 / 48.0, 3.0 / 48.0, 1.0 / 48.0},
	},
}
var burkes colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}
var stucki colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}
var atkinson colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}
var sierra3 colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}
var sierra2 colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}
var sierraLite colorquant.Dither = colorquant.Dither{
	Filter: [][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
		{4.0 / 32.0, 8.0 / 32.0, 0.0, 0.0, 0.0},
	},
}

var dither map[string]colorquant.Dither = map[string]colorquant.Dither{
	"FloydSteinberg": floydSteinberg,
	"Burkes":         burkes,
	"Stucki":         stucki,
	"Atkinson":       atkinson,
	"Sierra-3":       sierra3,
	"Sierra-2":       sierra2,
	"Sierra-Lite":    sierraLite,
}

var woolPalette []color.Color = []color.Color{
	color.RGBA{0xE9, 0xEC, 0xEC, 0xFF},
	color.RGBA{0xF0, 0x76, 0x13, 0xFF},
	color.RGBA{0xBD, 0x44, 0xB3, 0xFF},
	color.RGBA{0x3A, 0xAF, 0xD9, 0xFF},
	color.RGBA{0xF8, 0xC6, 0x27, 0xFF},
	color.RGBA{0x70, 0xB9, 0x19, 0xFF},
	color.RGBA{0xED, 0x8D, 0xAC, 0xFF},
	color.RGBA{0x3E, 0x44, 0x47, 0xFF},
	color.RGBA{0x8E, 0x8E, 0x86, 0xFF},
	color.RGBA{0x15, 0x89, 0x91, 0xFF},
	color.RGBA{0x79, 0x2A, 0xAC, 0xFF},
	color.RGBA{0x35, 0x39, 0x9D, 0xFF},
	color.RGBA{0x72, 0x47, 0x28, 0xFF},
	color.RGBA{0x54, 0x6D, 0x1B, 0xFF},
	color.RGBA{0xA1, 0x27, 0x22, 0xFF},
	color.RGBA{0x14, 0x15, 0x19, 0xFF},
	color.RGBA{0x00, 0x00, 0x00, 0x00},
}

var woolColorMap map[color.Color]string = map[color.Color]string{
	color.RGBA{0xE9, 0xEC, 0xEC, 0xFF}: "white_wool",
	color.RGBA{0xF0, 0x76, 0x13, 0xFF}: "orange_wool",
	color.RGBA{0xBD, 0x44, 0xB3, 0xFF}: "magenta_wool",
	color.RGBA{0x3A, 0xAF, 0xD9, 0xFF}: "light_blue_wool",
	color.RGBA{0xF8, 0xC6, 0x27, 0xFF}: "yellow_wool",
	color.RGBA{0x70, 0xB9, 0x19, 0xFF}: "lime_wool",
	color.RGBA{0xED, 0x8D, 0xAC, 0xFF}: "pink_wool",
	color.RGBA{0x3E, 0x44, 0x47, 0xFF}: "gray_wool",
	color.RGBA{0x8E, 0x8E, 0x86, 0xFF}: "light_gray_wool",
	color.RGBA{0x15, 0x89, 0x91, 0xFF}: "cyan_wool",
	color.RGBA{0x79, 0x2A, 0xAC, 0xFF}: "purple_wool",
	color.RGBA{0x35, 0x39, 0x9D, 0xFF}: "blue_wool",
	color.RGBA{0x72, 0x47, 0x28, 0xFF}: "brown_wool",
	color.RGBA{0x54, 0x6D, 0x1B, 0xFF}: "green_wool",
	color.RGBA{0xA1, 0x27, 0x22, 0xFF}: "red_wool",
	color.RGBA{0x14, 0x15, 0x19, 0xFF}: "black_wool",
	color.RGBA{0x00, 0x00, 0x00, 0x00}: "air",
}
