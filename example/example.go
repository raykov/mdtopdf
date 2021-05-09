package main

import (
	"fmt"
	"os"

	"github.com/raykov/mdtopdf"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("convert example.md example.pdf")
		return
	}
	mdPath := os.Args[1]
	pdfPath := os.Args[2]

	md, err := os.Open(mdPath)
	if err != nil{
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

	err = mdtopdf.Convert(md, pdf)
	if err != nil {
		fmt.Println(err)
		return
	}
}
