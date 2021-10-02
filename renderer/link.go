package renderer

import (
	"strings"

	"github.com/raykov/mdtopdf/color"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func renderAutoLink(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.AutoLink)
	if !entering {
		return ast.WalkContinue, nil
	}

	url := string(util.URLEscape(n.URL(source), false))
	if n.AutoLinkType == ast.AutoLinkEmail && !strings.HasPrefix(strings.ToLower(url), "mailto:") {
		w.SetLink("mailto:" + url)
	} else {
		w.SetLink(url)
	}

	w.Write(util.EscapeHTML(n.Label(source)))
	w.SetLink("")
	return ast.WalkContinue, nil
}

func renderLink(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		w.SetLink(string(util.URLEscape(n.Destination, true)))

		w.WithStyle(w.NewSetFontStyleOption(underscoreStyle), w.NewTextColorOption(color.Blue))
	} else {
		w.SetLink("")
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}
