package renderer

import "github.com/yuin/goldmark/ast"

func renderHTMLBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func renderRawHTML(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}
