package mdtopdf

import (
	"io"

	"github.com/raykov/mdtopdf/document"
	"github.com/raykov/mdtopdf/renderer"

	"github.com/raykov/gofpdf"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Option customises Convert.
type Option func(*config)

type config struct {
	extensions []func(*gofpdf.Fpdf)
	inline     document.InlineImager
}

// WithExtensions applies the functions to the PDF before any content is added.
func WithExtensions(extensions ...func(*gofpdf.Fpdf)) Option {
	return func(c *config) { c.extensions = append(c.extensions, extensions...) }
}

// WithInlineImager lets i draw parts of the text as pictures, e.g. color emoji.
func WithInlineImager(i document.InlineImager) Option {
	return func(c *config) { c.inline = i }
}

// Convert your Markdown to PDF
func Convert(r io.Reader, w io.Writer, extensions ...func(*gofpdf.Fpdf)) error {
	return ConvertWithOptions(r, w, WithExtensions(extensions...))
}

// ConvertWithOptions converts your Markdown to PDF, applying the options
func ConvertWithOptions(r io.Reader, w io.Writer, opts ...Option) error {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	md, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	markdown := goldmark.New(
		goldmark.WithRenderer(renderer.NewRenderer()),
		goldmark.WithExtensions(
			extension.NewTable(),
			extension.Strikethrough,
		),
	)

	pdf := gofpdf.New("P", "pt", "A4", ".")

	for _, extension := range cfg.extensions {
		extension(pdf)
	}

	pdf.AddPage()

	d := document.NewDocument(pdf, document.DefaultStyle)
	d.SetInlineImager(cfg.inline)

	if err = markdown.Convert(md, d); err != nil {
		return err
	}

	return pdf.Output(w)
}
