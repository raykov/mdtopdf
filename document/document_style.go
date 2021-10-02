package document

func (d *Document) WithStyle(opts ...func(style interface{}) error) {
	newStyle := d.styles.Current().Clone()

	for _, opt := range opts {
		_ = opt(newStyle)
	}

	d.withStyle(newStyle)
}

// ResetStyle resets to previous style in PDF
func (d *Document) ResetStyle() {
	_ = d.styles.Pop()
	d.setStyle(d.styles.Current())
}

func (d *Document) withStyle(s *Style) {
	if s.LeftMargin == 0 {
		s.LeftMargin, _, _, _ = d.PDF.GetMargins()
		s.CellMargin = s.LeftMargin / 10
	}

	d.styles.Add(s)
	d.setStyle(s)
}

// setStyle sets style configs in PDF
func (d *Document) setStyle(s *Style) {
	d.PDF.SetFont(s.FontFamily, s.FontStyle, s.FontSize)
	d.PDF.SetTextColor(s.TextColor.R, s.TextColor.G, s.TextColor.B)
	d.PDF.SetFillColor(s.FillColor.R, s.FillColor.G, s.FillColor.B)
	d.PDF.SetDrawColor(1, 1, 1)
	d.PDF.SetLeftMargin(s.LeftMargin)
	d.PDF.SetCellMargin(s.CellMargin)
	//d.PDF.SetLineCapStyle("round")
	//d.PDF.SetTextRenderingMode(1)
	//d.PDF.CellFormat(1000, 20, "lipsum", "1", 0, "", true, 0, "")
	//d.PDF.MultiCell(1000, 15, "txtStr", "1", "", true)
}
