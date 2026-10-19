package document

import (
	"bytes"

	"github.com/raykov/gofpdf"
)

const (
	// inlineImageGap is the space after an inline image, relative to its size.
	inlineImageGap = 0.08
	// tableImageBaseline is how far (relative to the size) an inline image in a
	// table cell reaches above the text baseline.
	tableImageBaseline = 0.82
)

// InlineImage is a picture drawn in the flow of a text, like an emoji.
type InlineImage struct {
	// Name identifies the image in the PDF; images with the same name are embedded once.
	Name string
	// Type is the image format, as understood by gofpdf: "png", "jpg", "gif"...
	Type string
	// Data holds the encoded image.
	Data []byte
}

// Segment is a part of a text: either a plain string, or, when Image is set,
// a picture drawn instead of Text.
type Segment struct {
	Text  string
	Image *InlineImage
}

// InlineImager splits a text into plain parts and parts drawn as pictures.
// It lets callers render characters that the PDF fonts cannot, e.g. color emoji.
type InlineImager interface {
	Segments(text string) []Segment
}

// SetInlineImager sets (or, with nil, removes) the InlineImager used for all text.
func (d *Document) SetInlineImager(i InlineImager) {
	d.inline = i
}

func (d *Document) segments(s string) []Segment {
	if d.inline == nil {
		return []Segment{{Text: s}}
	}
	return d.inline.Segments(s)
}

// writeInlineImage draws the image at the cursor, wrapping to a new line (or
// page) when it does not fit, and moves the cursor after it.
func (d *Document) writeInlineImage(lineHeight, size float64, img *InlineImage, link string) {
	pageW, pageH := d.PDF.GetPageSize()
	_, _, right, _ := d.PDF.GetMargins()

	// gofpdf draws the text of Write one cell margin to the right of the cursor
	margin := d.PDF.GetCellMargin()

	x, y := d.PDF.GetXY()
	if x+margin+size > pageW-right+0.01 {
		d.PDF.Ln(lineHeight)
		x, y = d.PDF.GetXY()
	}
	if auto, bottom := d.PDF.GetAutoPageBreak(); auto && y+lineHeight > pageH-bottom {
		d.PDF.AddPage()
		x, y = d.PDF.GetXY()
	}

	opts := d.registerInlineImage(img)
	d.PDF.ImageOptions(img.Name, x+margin, y+(lineHeight-size)/2, size, size, false, opts, 0, link)
	d.PDF.SetX(x + size*(1+inlineImageGap))
}

// text draws s with its baseline at (x, y), like gofpdf's Text, but with
// inline images for the parts the InlineImager asks for.
func (d *Document) text(x, y float64, s string) {
	size := d.styles.Current().FontSize
	for _, seg := range d.segments(s) {
		if seg.Image == nil {
			d.PDF.Text(x, y, seg.Text)
			x += d.PDF.GetStringWidth(seg.Text)
			continue
		}
		opts := d.registerInlineImage(seg.Image)
		d.PDF.ImageOptions(seg.Image.Name, x, y-size*tableImageBaseline, size, size, false, opts, 0, "")
		x += size * (1 + inlineImageGap)
	}
}

func (d *Document) registerInlineImage(img *InlineImage) gofpdf.ImageOptions {
	opts := gofpdf.ImageOptions{ImageType: img.Type}
	d.PDF.RegisterImageOptionsReader(img.Name, opts, bytes.NewReader(img.Data))
	return opts
}
