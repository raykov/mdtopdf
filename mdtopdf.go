package mdtopdf

import (
	"io"

	"github.com/Acke0/mdtopdf/document"
	"github.com/Acke0/mdtopdf/renderer"

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

	style := SetStyleFromPdf(pdf)
	d := document.NewDocument(pdf, style)

	if err = markdown.Convert(md, d); err != nil {
		return err
	}

	err = pdf.Output(w)
	if err != nil {
		return err
	}

	return nil
}

// Use styles defined in PDF extensions
func SetStyleFromPdf(pdf *gofpdf.Fpdf) (style *document.Style) {
	style = document.DefaultStyle

	style.FontSize, _ = pdf.GetFontSize()
	style.FillColor.R, style.FillColor.G, style.FillColor.B = pdf.GetFillColor()
	style.TextColor.R, style.TextColor.G, style.TextColor.B = pdf.GetTextColor()
	style.LeftMargin, _, _, _ = pdf.GetMargins()
	style.CellMargin = pdf.GetCellMargin()

	return
}
