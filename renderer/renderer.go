package renderer

import (
	"io"

	"github.com/yuin/goldmark/ast"
	tast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
)

// NewRenderer returns new default renderer
func NewRenderer() *Renderer {
	r := &Renderer{}

	// table
	r.Register(tast.KindTable, renderTable)
	r.Register(tast.KindTableHeader, renderTableHeader)
	r.Register(tast.KindTableRow, renderTableRow)
	r.Register(tast.KindTableCell, renderTableCell)
	// blocks
	r.Register(ast.KindDocument, renderDocument)
	r.Register(ast.KindHeading, renderHeading)
	r.Register(ast.KindBlockquote, renderBlockquote)
	r.Register(ast.KindCodeBlock, renderCodeBlock)
	r.Register(ast.KindFencedCodeBlock, renderFencedCodeBlock)
	r.Register(ast.KindHTMLBlock, renderHTMLBlock)
	r.Register(ast.KindList, renderList)
	r.Register(ast.KindListItem, renderListItem)
	r.Register(ast.KindParagraph, renderParagraph)
	r.Register(ast.KindTextBlock, renderTextBlock)
	r.Register(ast.KindThematicBreak, renderThematicBreak)
	// inlines
	r.Register(ast.KindAutoLink, renderAutoLink)
	r.Register(ast.KindCodeSpan, renderCodeSpan)
	r.Register(ast.KindEmphasis, renderEmphasis)
	r.Register(ast.KindImage, renderImage)
	r.Register(ast.KindLink, renderLink)
	r.Register(ast.KindRawHTML, renderRawHTML)
	r.Register(ast.KindText, renderText)
	r.Register(ast.KindString, renderString)
	// strikethrough
	r.Register(tast.KindStrikethrough, renderStrikethrough)

	return r
}

// Renderer holds all functions to render
type Renderer struct {
	nodeRendererFuncs map[ast.NodeKind]NodeRendererFunc
}

// Render renders a PDF doc
func (r *Renderer) Render(w io.Writer, source []byte, node ast.Node) error {
	writer := w.(any).(Writer)

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
