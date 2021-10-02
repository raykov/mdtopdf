package renderer

import (
	"github.com/yuin/goldmark/ast"
)

const (
	italicStyle     = "I"
	boldStyle       = "B"
	underscoreStyle = "U"
	strikeStyle     = "S"
)

func renderEmphasis(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	//"B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)
	n := node.(*ast.Emphasis)
	style := italicStyle
	if n.Level == 2 {
		style = boldStyle
	}

	if entering {
		w.WithStyle(w.NewAddFontStyleOption(style))
	} else {
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}

func renderStrikethrough(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	//"B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)

	if entering {
		w.WithStyle(w.NewAddFontStyleOption(strikeStyle))
	} else {
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}
