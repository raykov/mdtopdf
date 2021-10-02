package renderer

import (
	"github.com/yuin/goldmark/ast"
)

func renderParagraph(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.Write([]byte("\n\n"))
	}

	return ast.WalkContinue, nil
}

func renderTextBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if _, ok := node.NextSibling().(ast.Node); ok && node.FirstChild() != nil {
			_, _ = w.Write([]byte("\n"))
		}
	}
	return ast.WalkContinue, nil
}

func renderText(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	segment := n.Segment

	if n.IsRaw() {
		_, _ = w.Write(segment.Value(source))
	} else {
		_, _ = w.Write(segment.Value(source))
		if n.HardLineBreak() { // || (n.SoftLineBreak() && r.HardWraps)
			_, _ = w.Write([]byte("\n\n"))
		} else if n.SoftLineBreak() {
			_, _ = w.Write([]byte("\n"))
		}
	}

	return ast.WalkContinue, nil
}

func renderString(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.String)

	if n.IsCode() {
		_, _ = w.Write(n.Value)
	} else {
		if n.IsRaw() {
			_, _ = w.Write(n.Value)
		} else {
			_, _ = w.Write(n.Value)
		}
	}
	return ast.WalkContinue, nil
}
