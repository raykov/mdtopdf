# mdtopdf
Converts Markdown to PDF


```go
	md, err := os.Open("example.md")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer md.Close()

	pdf, err := os.Create("example.pdf")
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

	err = mdtopdf.Convert(md, pdf, pageNumExtension)
	if err != nil {
		fmt.Println(err)
		return
	}
```