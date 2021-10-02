package renderer

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
)

func renderList(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WithStyle(w.NewAddMarginOption(10))
	} else {
		w.ResetStyle()
		_, _ = w.Write([]byte("\n"))
	}
	return ast.WalkContinue, nil
}

func renderListItem(w Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if p := n.Parent().(*ast.List); p.IsOrdered() {
			s := p.Start
			i := 0
			for c := p.FirstChild(); c != nil; i++ {
				if c == n {
					break
				}
				c = c.NextSibling()
			}

			_, _ = w.Write([]byte(fmt.Sprintf("%d. ", s+i)))
		} else {
			_, _ = w.Write([]byte("* "))
		}
	} else {
		_, _ = w.Write([]byte("\n"))
	}
	return ast.WalkContinue, nil
}
