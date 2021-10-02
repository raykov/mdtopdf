package renderer

import "github.com/yuin/goldmark/ast"

func renderThematicBreak(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	w.ThematicBreak()

	return ast.WalkContinue, nil
}
