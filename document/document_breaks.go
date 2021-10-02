package document

// LineBreak goes to the beginning of the next line
func (d *Document) LineBreak() {
	d.PDF.Ln(-1)
}

// ThematicBreak add a line break
func (d *Document) ThematicBreak() {
	dw, _ := d.PDF.GetPageSize()
	lm, _, rm, _ := d.PDF.GetMargins()

	w := dw - lm - rm
	h := d.styles.Current().FontSize

	d.PDF.SetY(d.PDF.GetY() + h)
	d.PDF.CellFormat(w, h, "", "T", 0, "", true, 0, "")
	d.PDF.Ln(-1)
}
