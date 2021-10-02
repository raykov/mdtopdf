package renderer

import (
	"github.com/raykov/mdtopdf/color"

	"github.com/yuin/goldmark/ast"
)

var headingDefaults = map[int]*HeadingStyle{
	1: {FontFamily: "Arial", FontStyle: "", FontSize: 24, Padding: 5, TextColor: color.Black, FillColor: color.White},
	2: {FontFamily: "Arial", FontStyle: "", FontSize: 22, Padding: 5, TextColor: color.Black, FillColor: color.White},
	3: {FontFamily: "Arial", FontStyle: "", FontSize: 20, Padding: 5, TextColor: color.Black, FillColor: color.White},
	4: {FontFamily: "Arial", FontStyle: "", FontSize: 18, Padding: 5, TextColor: color.Black, FillColor: color.White},
	5: {FontFamily: "Arial", FontStyle: "", FontSize: 16, Padding: 5, TextColor: color.Black, FillColor: color.White},
	6: {FontFamily: "Arial", FontStyle: "", FontSize: 14, Padding: 5, TextColor: color.Black, FillColor: color.White},
}

type HeadingStyle struct {
	FontFamily string
	FontStyle  string
	FontSize   float64

	TextColor *color.Color
	FillColor *color.Color

	Padding    float64
	LeftMargin float64
}

func renderHeading(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)
	if entering {
		h := headingDefaults[n.Level]
		w.WithStyle(
			w.NewFontFamilyOption(h.FontFamily),
			w.NewSetFontStyleOption(h.FontStyle),
			w.NewFontSizeOption(h.FontSize),
		)
	} else {
		_, _ = w.Write([]byte("\n"))
		w.ResetStyle()
	}

	return ast.WalkContinue, nil
}
