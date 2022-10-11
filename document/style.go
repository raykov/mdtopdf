package document

import "github.com/Acke0/mdtopdf/color"

var DefaultStyle = &Style{FontFamily: "Arial", FontStyle: "", FontSize: 12, TextColor: color.Black, FillColor: color.White, Padding: 2, CellMargin: 1}
var DefaultStyleBold = &Style{FontFamily: "Arial", FontStyle: "B", FontSize: 12, TextColor: color.Black, FillColor: color.White, Padding: 2, CellMargin: 1}
var DefaultStyleItalic = &Style{FontFamily: "Arial", FontStyle: "I", FontSize: 12, TextColor: color.Black, FillColor: color.White, Padding: 2, CellMargin: 1}

// Style for styling texts on the doc
type Style struct {
	FontFamily string
	FontStyle  string
	FontSize   float64

	TextColor *color.Color
	FillColor *color.Color

	Padding    float64
	LeftMargin float64
	CellMargin float64
}

// Clone creates a clone of the style
func (s *Style) Clone() *Style {
	return &Style{
		FontFamily: s.FontFamily,
		FontStyle:  s.FontStyle,
		FontSize:   s.FontSize,
		TextColor:  s.TextColor,
		FillColor:  s.FillColor,
		Padding:    s.Padding,
		LeftMargin: s.LeftMargin,
		CellMargin: s.CellMargin,
	}
}
