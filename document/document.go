package document

import (
	"github.com/raykov/mdtopdf/color"

	"github.com/raykov/gofpdf"
)

// NewDocument creates a new document
func NewDocument(pdf *gofpdf.Fpdf, defaultStyle *Style) *Document {
	styles := &Styles{}

	d := &Document{
		PDF:    pdf,
		styles: styles,
	}

	d.withStyle(defaultStyle)

	pdf.SetDrawColor(color.Black.R, color.Black.G, color.Black.B)
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
