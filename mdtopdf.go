package mdtopdf

import (
	"io"

	"github.com/raykov/mdtopdf/document"
	"github.com/raykov/mdtopdf/renderer"

	"github.com/raykov/gofpdf"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Convert your Markdown to PDF
func Convert(r io.Reader, w io.Writer, extensions ...func(*gofpdf.Fpdf)) error {
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

	for _, extension := range extensions {
		extension(pdf)
	}

	pdf.AddPage()

	d := document.NewDocument(pdf, document.DefaultStyle)

	if err = markdown.Convert(md, d); err != nil {
		return err
	}

	err = pdf.Output(w)
	if err != nil {
		return err
	}

	return nil
}
