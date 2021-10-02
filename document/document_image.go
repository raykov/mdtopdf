package document

import (
	"fmt"
	"io"
	"strings"

	"github.com/raykov/gofpdf"

	"github.com/yuin/goldmark/ast"
)

// WriteImage draw an image in PDF
func (d *Document) WriteImage(path string, file io.Reader, _ ...ast.Attribute) {
	width, _ := d.PDF.GetPageSize()
	left, _, right, _ := d.PDF.GetMargins()
	pw := width - left - right

	options := gofpdf.ImageOptions{}

	if options.ImageType == "" {
		pos := strings.LastIndex(path, ".")
		if pos < 0 {
			fmt.Printf("image file has no extension and no type was specified: %s", path)
			return
		}
		options.ImageType = path[pos+1:]
	}

	info := d.PDF.RegisterImageOptionsReader(path, options, file)

	w := info.Width()
	if w > pw {
		w = pw
	}

	d.PDF.ImageOptions(path,
		-1, 0, w, 0, true,
		gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
}
