package main

import (
	"fmt"
	"os"

	"github.com/raykov/gofpdf"

	"github.com/raykov/mdtopdf"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("convert example.md example.pdf")
		return
	}
	mdPath := os.Args[1]
	pdfPath := os.Args[2]

	md, err := os.Open(mdPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer md.Close()

	pdf, err := os.Create(pdfPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdf.Close()

	pageNumExtension := func(pdf *gofpdf.Fpdf) {
		pdf.SetFooterFunc(func() {
			left, _, right, bottom := pdf.GetMargins()
			width, height := pdf.GetPageSize()
			fontSize := 12.0

			pNum := fmt.Sprint(pdf.PageNo())
			pdf.SetXY(width-left/2-pdf.GetStringWidth(pNum), height-bottom/2)
			pdf.SetFontSize(fontSize)
			pdf.SetTextColor(200, 200, 200)
			pdf.SetFontStyle("B")
			pdf.SetRightMargin(0)
			pdf.Write(fontSize, pNum)
			pdf.SetRightMargin(right)
		})
	}

	err = mdtopdf.Convert(
		md,
		pdf,
		pageNumExtension,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
