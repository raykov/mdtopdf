package document

import (
	"github.com/raykov/mdtopdf/color"
)

// WriteTable draw a table in PDF
func (d *Document) WriteTable(rows ...[]string) {
	if len(rows) == 0 {
		return
	}
	currStyle := d.styles.Current()
	hStyle := currStyle.Clone()
	hStyle.FontStyle += "B"
	hStyle.FillColor = color.LightGray

	ocw := d.PDF.GetStringWidth("W")

	dw, dh := d.PDF.GetPageSize()
	lm, tm, rm, bm := d.PDF.GetMargins()

	w := dw - lm - rm
	h := dh - bm

	cw := w / float64(len(rows[0]))
	ch := currStyle.FontSize + currStyle.Padding + 2

	for ri, row := range rows {
		if ri == 0 {
			d.withStyle(hStyle)
		}
		rLines := make([][][]byte, 0, len(row))
		maxLinesNum := 0
		for _, s := range row {
			lines := d.PDF.SplitLines([]byte(s), cw)
			if len(s) == 0 {
				lines = [][]byte{{' '}}
			}
			if l := len(lines); l > maxLinesNum {
				maxLinesNum = l
			}
			rLines = append(rLines, lines)
		}
		x, y := d.PDF.GetXY()
		sp := d.PDF.PageNo()
		d.PDF.CellFormat(w, ch*float64(maxLinesNum), "", "1", 0, "", true, 0, "")
		ep := d.PDF.PageNo()

		for i := range row {
			if ep > sp {
				d.PDF.SetPage(sp)
			}
			np := false
			for j, line := range rLines[i] {
				tx := x + cw*float64(i) + ocw/4 + 3
				ty := y + ch*float64(j) + currStyle.FontSize
				if ep > sp && ty+ch > h {
					d.PDF.SetPage(ep)
					np = true
				}
				if np {
					ty = tm + ch*float64(j) + currStyle.FontSize
				}

				d.PDF.Text(tx, ty, string(line))
			}
		}

		if ri == 0 {
			d.ResetStyle()
		}
		d.LineBreak()
	}
	d.PDF.Ln(1)
}
