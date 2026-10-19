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
	inline          InlineImager
}

// Write writes a simple slice of bytes.
// Also needed to satisfy io.Writer interface for `goldmark`
func (d *Document) Write(p []byte) (n int, err error) {
	currStyle := d.styles.Current()
	lineHeight := currStyle.FontSize + currStyle.Padding

	for _, seg := range d.segments(string(p)) {
		switch {
		case seg.Image != nil:
			d.writeInlineImage(lineHeight, currStyle.FontSize, seg.Image, d.linkDestination)
		case d.linkDestination != "":
			d.PDF.WriteLinkString(lineHeight, seg.Text, d.linkDestination)
		default:
			d.PDF.Write(lineHeight, seg.Text)
		}
	}

	return len(p), nil
}

// SetLink set a variable to specify link element
func (d *Document) SetLink(link string) {
	d.linkDestination = link
}
