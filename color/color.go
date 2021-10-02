package color

var Black = &Color{0, 0, 0}
var Blue = &Color{100, 100, 255}
var LightGray = &Color{200, 200, 200}
var VeryLightGray = &Color{230, 230, 230}
var White = &Color{255, 255, 255}

// Color just a color
type Color struct {
	R, G, B int
}
