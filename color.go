package mdtopdf

var black = &Color{0, 0, 0}
var white = &Color{255, 255, 255}
var blue = &Color{100, 100, 255}
var lightGray = &Color{200, 200, 200}

// Color just a color
type Color struct {
	R, G, B int
}
