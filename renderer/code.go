package renderer

import (
	"github.com/raykov/mdtopdf/color"

	"github.com/yuin/goldmark/ast"
)

func renderCodeBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WithStyle(w.NewFillColorOption(color.VeryLightGray))
	} else {
		w.ResetStyle()
	}

	return ast.WalkContinue, nil
}

func renderCodeSpan(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WithStyle(w.NewFillColorOption(color.VeryLightGray))
	} else {
		w.ResetStyle()
	}

	return ast.WalkContinue, nil
}

func renderFencedCodeBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	cb := node.(*ast.FencedCodeBlock)

	if entering {
		w.WithStyle(w.NewFillColorOption(color.VeryLightGray), w.NewAddCellMarginOption(10))

		_, _ = w.Write([]byte("\n"))
		ls := cb.Lines()
		for i := 0; i < ls.Len(); i++ {
			seg := ls.At(i)

			_, _ = w.Write(seg.Value(source))
		}
		_, _ = w.Write([]byte("\n"))
	} else {
		w.ResetStyle()
		_, _ = w.Write([]byte("\n"))
	}

	return ast.WalkContinue, nil
}
