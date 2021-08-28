package mdtopdf

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/yuin/goldmark/ast"
	tast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// NewRenderer returns new default renderer
func NewRenderer() *Renderer {
	r := &Renderer{}

	// table
	r.Register(tast.KindTable, r.renderTable)
	r.Register(tast.KindTableHeader, r.renderTableHeader)
	r.Register(tast.KindTableRow, r.renderTableRow)
	r.Register(tast.KindTableCell, r.renderTableCell)
	// blocks
	r.Register(ast.KindDocument, r.renderDocument)
	r.Register(ast.KindHeading, r.renderHeading)
	r.Register(ast.KindBlockquote, r.renderBlockquote)
	r.Register(ast.KindCodeBlock, r.renderCodeBlock)
	r.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	r.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	r.Register(ast.KindList, r.renderList)
	r.Register(ast.KindListItem, r.renderListItem)
	r.Register(ast.KindParagraph, r.renderParagraph)
	r.Register(ast.KindTextBlock, r.renderTextBlock)
	r.Register(ast.KindThematicBreak, r.renderThematicBreak)
	// inlines
	r.Register(ast.KindAutoLink, r.renderAutoLink)
	r.Register(ast.KindCodeSpan, r.renderCodeSpan)
	r.Register(ast.KindEmphasis, r.renderEmphasis)
	r.Register(ast.KindImage, r.renderImage)
	r.Register(ast.KindLink, r.renderLink)
	r.Register(ast.KindRawHTML, r.renderRawHTML)
	r.Register(ast.KindText, r.renderText)
	r.Register(ast.KindString, r.renderString)
	// strikethrough
	r.Register(tast.KindStrikethrough, r.renderStrikethrough)

	return r
}

// Renderer holds all functions to render
type Renderer struct {
	nodeRendererFuncs map[ast.NodeKind]NodeRendererFunc
}

// Render renders a PDF doc
func (r *Renderer) Render(w io.Writer, source []byte, node ast.Node) error {

	writer := w.(interface{}).(Writer)

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		var err error
		f := r.nodeRendererFuncs[node.Kind()]

		if f != nil {
			s, err = f(writer, source, node, entering)
		}
		return s, err
	})

	if err != nil {
		return err
	}

	return nil
}

// AddOptions adds given option to this renderer.
func (r *Renderer) AddOptions(...renderer.Option) {}

// NodeRendererFunc is a function that renders a given node.
type NodeRendererFunc func(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error)

// Register a new node render func
func (r *Renderer) Register(kind ast.NodeKind, v NodeRendererFunc) {
	if r.nodeRendererFuncs == nil {
		r.nodeRendererFuncs = map[ast.NodeKind]NodeRendererFunc{}
	}

	r.nodeRendererFuncs[kind] = v
}

// render Funcs

func (r *Renderer) renderTable(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		data := make([][]string, 0, node.ChildCount())
		for row := node.FirstChild(); row != nil; {
			subData := make([]string, 0, row.ChildCount())

			for cell := row.FirstChild(); cell != nil; {
				subData = append(subData, string(cell.Text(source)))

				cell = cell.NextSibling()
			}
			data = append(data, subData)
			row = row.NextSibling()
		}

		w.WriteTable(data...)

		return ast.WalkSkipChildren, nil
	}

	_, _ = w.Write([]byte("\n"))

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTableHeader(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTableRow(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderTableCell(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderDocument(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)
	if entering {
		w.WithStyle(H[n.Level])
	} else {
		_, _ = w.Write([]byte("\n"))
		w.ResetStyle()
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderBlockquote(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderFencedCodeBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		currStyle := w.CurrentStyle()
		newStyle := currStyle.Clone()
		newStyle.LeftMargin += 10
		w.WithStyle(newStyle)
	} else {
		w.ResetStyle()
		_, _ = w.Write([]byte("\n"))
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderParagraph(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
	} else {
		_, _ = w.Write([]byte("\n\n"))
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if _, ok := node.NextSibling().(ast.Node); ok && node.FirstChild() != nil {
			_, _ = w.Write([]byte("\n"))
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	w.ThematicBreak()
	return ast.WalkContinue, nil
}

func (r *Renderer) renderAutoLink(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderCodeSpan(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderEmphasis(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	//"B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)
	n := node.(*ast.Emphasis)
	style := "I"
	if n.Level == 2 {
		style = "B"
	}

	if entering {
		currStyle := w.CurrentStyle()
		newStyle := currStyle.Clone()
		newStyle.FontStyle += style
		w.WithStyle(newStyle)
	} else {
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Image)

	filePath := string(util.URLEscape(n.Destination, true))
	inlineImg := "data:image/"

	switch {
	case strings.HasPrefix(filePath, inlineImg):
		parts := strings.Split(filePath, ",")
		if len(parts) != 2 {
			return ast.WalkSkipChildren, nil
		}
		partsInfo := strings.Split(strings.TrimPrefix(parts[0], inlineImg), ";")
		if len(partsInfo) != 2 {
			return ast.WalkSkipChildren, nil
		}
		if partsInfo[1] != "base64" {
			return ast.WalkSkipChildren, nil
		}
		data, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return ast.WalkSkipChildren, nil
		}

		file := bytes.NewReader(data)
		w.WriteImage(parts[1]+"."+partsInfo[0], file, node.Attributes()...)
	case strings.HasPrefix(filePath, "http"):
		resp, err := http.Get(filePath)
		if err != nil {
			fmt.Printf("File NOT found: %s\n", filePath)
			return ast.WalkSkipChildren, nil
		}
		defer resp.Body.Close()

		w.WriteImage(filePath, resp.Body, node.Attributes()...)
	default:
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("File NOT found: %s\n", filePath)
			return ast.WalkSkipChildren, nil
		}
		defer file.Close()
		w.WriteImage(filePath, file, node.Attributes()...)
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderLink(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		w.SetLink(string(util.URLEscape(n.Destination, true)))
		currStyle := w.CurrentStyle()
		newStyle := currStyle.Clone()
		newStyle.FontStyle += "U"
		newStyle.TextColor = blue
		w.WithStyle(newStyle)
	} else {
		w.SetLink("")
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderRawHTML(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderString(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderStrikethrough(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	//"B" (bold), "I" (italic), "U" (underscore), "S" (strike-out)

	if entering {
		currStyle := w.CurrentStyle()
		newStyle := currStyle.Clone()
		newStyle.FontStyle += "S"
		w.WithStyle(newStyle)
	} else {
		w.ResetStyle()
	}
	return ast.WalkContinue, nil
}
