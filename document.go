package mdtopdf

import (
	"fmt"
	"io"
	"strings"

	"github.com/jung-kurt/gofpdf"

	"github.com/yuin/goldmark/ast"
)

// NewDocument creates a new document
func NewDocument(pdf *gofpdf.Fpdf, defaultStyle *Style) *Document {
	styles := &Styles{}

	d := &Document{
		PDF:    pdf,
		styles: styles,
	}

	d.WithStyle(defaultStyle)

	pdf.SetDrawColor(black.R, black.G, black.B)
	pdf.SetLineWidth(.5)

	return d
}

// Document a layer between Markdown and PDF
type Document struct {
	PDF    *gofpdf.Fpdf
	styles *Styles

	linkDestination string
}

// Write writes a simple slice of bytes.
// Also needed to satisfy io.Writer interface for `goldmark`
func (d *Document) Write(p []byte) (n int, err error) {
	s := string(p)

	currStyle := d.styles.Current()

	switch {
	default:
		d.PDF.Write(currStyle.FontSize+currStyle.Padding, s)
	case d.linkDestination != "":
		d.PDF.WriteLinkString(currStyle.FontSize+currStyle.Padding, s, d.linkDestination)
	}

	return len(p), nil
}

// SetLink set a variable to specify link element
func (d *Document) SetLink(link string) {
	d.linkDestination = link
}

// WriteTable draw a table in PDF
func (d *Document) WriteTable(rows ...[]string) {
	if len(rows) == 0 {
		return
	}
	currStyle := d.CurrentStyle()
	hStyle := currStyle.Clone()
	hStyle.FontStyle += "B"
	hStyle.FillColor = lightGray

	ocw := d.PDF.GetStringWidth("W")

	dw, dh := d.PDF.GetPageSize()
	lm, tm, rm, bm := d.PDF.GetMargins()

	w := dw - lm - rm
	h := dh - bm

	cw := w / float64(len(rows[0]))
	ch := currStyle.FontSize + currStyle.Padding + 2

	for ri, row := range rows {
		if ri == 0 {
			d.WithStyle(hStyle)
		}
		rLines := make([][][]byte, 0, len(row))
		maxLinesNum := 0
		for _, s := range row {
			lines := d.PDF.SplitLines([]byte(s), cw)
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
				tx := x + cw*float64(i) + ocw/4
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
}

// LineBreak goes to the beginning of the next line
func (d *Document) LineBreak() {
	d.PDF.Ln(-1)
}

// ThematicBreak add a line break
func (d *Document) ThematicBreak() {
	dw, _ := d.PDF.GetPageSize()
	lm, _, rm, _ := d.PDF.GetMargins()

	w := dw - lm - rm
	h := d.CurrentStyle().FontSize

	d.PDF.SetY(d.PDF.GetY() + h)
	d.PDF.CellFormat(w, h, "", "T", 0, "", true, 0, "")
	d.PDF.Ln(-1)
}

// WriteImage draw an image in PDF
func (d *Document) WriteImage(path string, file io.Reader, _ ...ast.Attribute) {
	width, _ := d.PDF.GetPageSize()
	left, _, right, _ := d.PDF.GetMargins()
	pw := width - left - right

	options := gofpdf.ImageOptions{}

	if options.ImageType == "" {
		pos := strings.LastIndex(path, ".")
		if pos < 0 {
			fmt.Errorf("image file has no extension and no type was specified: %s", path)
			return
		}
		options.ImageType = path[pos+1:]
	}

	info := d.PDF.RegisterImageOptionsReader(path, options, file)

	w := info.Width()
	if w > pw {
		w = pw
	}

	d.PDF.ImageOptions(path,
		-1, 0, w, 0, true,
		gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
}

// WithStyle applies a new style to the PDF
func (d *Document) WithStyle(s *Style) Writer {
	if s.LeftMargin == 0 {
		s.LeftMargin, _, _, _ = d.PDF.GetMargins()
	}

	d.styles.Add(s)
	d.setStyle(s)

	return d
}

// ResetStyle resets to previous style in PDF
func (d *Document) ResetStyle() Writer {
	_ = d.styles.Pop()
	d.setStyle(d.styles.Current())

	return d
}

// CurrentStyle returns a current style of the PDF
func (d *Document) CurrentStyle() *Style {
	return d.styles.Current()
}

// setStyle sets style configs in PDF
func (d *Document) setStyle(s *Style) {
	d.PDF.SetFont(s.FontFamily, s.FontStyle, s.FontSize)
	d.PDF.SetTextColor(s.TextColor.R, s.TextColor.G, s.TextColor.B)
	d.PDF.SetFillColor(s.FillColor.R, s.FillColor.G, s.FillColor.B)
	d.PDF.SetLeftMargin(s.LeftMargin)
}
