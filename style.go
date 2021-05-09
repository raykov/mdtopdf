package mdtopdf

//"B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)
var defaultStyle = &Style{FontFamily: "Arial", FontStyle: "", FontSize: 12, TextColor: black, FillColor: white, Padding: 2}
var defaultStyleBold = &Style{FontFamily: "Arial", FontStyle: "B", FontSize: 12, TextColor: black, FillColor: white, Padding: 2}
var defaultStyleItalic = &Style{FontFamily: "Arial", FontStyle: "I", FontSize: 12, TextColor: black, FillColor: white, Padding: 2}

// H for Heading
var H = map[int]*Style{
	1: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 24, Padding: 5, TextColor: black, FillColor: white},
	2: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 22, Padding: 5, TextColor: black, FillColor: white},
	3: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 20, Padding: 5, TextColor: black, FillColor: white},
	4: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 18, Padding: 5, TextColor: black, FillColor: white},
	5: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 16, Padding: 5, TextColor: black, FillColor: white},
	6: &Style{FontFamily: "Arial", FontStyle: "", FontSize: 14, Padding: 5, TextColor: black, FillColor: white},
}

// Style for styling texts on the doc
type Style struct {
	FontFamily string
	FontStyle  string
	FontSize   float64

	TextColor *Color
	FillColor *Color

	Padding    float64
	LeftMargin float64
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
	}
}
