package mdtopdf

import (
	"io"

	"github.com/jung-kurt/gofpdf"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Convert your Markdown to PDF
func Convert(r io.Reader, w io.Writer) error {
	md, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	markdown := goldmark.New(
		goldmark.WithRenderer(NewRenderer()),
		goldmark.WithExtensions(
			extension.NewTable(),
		),
	)

	pdf := gofpdf.New("portrait", "pt", "Letter", ".")
	pdf.AddPage()

	d := NewDocument(pdf, defaultStyle)

	if err = markdown.Convert(md, d); err != nil {
		return err
	}

	err = pdf.Output(w)
	if err != nil {
		return err
	}

	return nil
}




