# mdtopdf
Converts Markdown to PDF


```go
	md, err := os.Open('example.md')
	if err != nil{
		fmt.Println(err)
		return
	}
	defer md.Close()

	pdf, err := os.Create('example.pdf')
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdf.Close()

	err = mdtopdf.Convert(md, pdf)
	if err != nil {
		fmt.Println(err)
		return
	}
```