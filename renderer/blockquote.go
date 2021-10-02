package renderer

import (
	"github.com/yuin/goldmark/ast"
)

func renderBlockquote(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WithStyle(w.NewAddFontStyleOption(italicStyle), w.NewFontSizeOption(10), w.NewFontFamilyOption("courier"))
	} else {
		w.ResetStyle()
	}

	return ast.WalkContinue, nil
}
