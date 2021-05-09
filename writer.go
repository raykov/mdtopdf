package mdtopdf

import (
	"io"

	"github.com/yuin/goldmark/ast"
)

// Writer describes what kind of behaviour we expect from writer
type Writer interface {
	Write(p []byte) (n int, err error)
	WithStyle(*Style) Writer
	ResetStyle() Writer
	CurrentStyle() *Style
	WriteImage(string, io.Reader, ...ast.Attribute)
	SetLink(string)
	WriteTable(rows ...[]string)
	LineBreak()
	ThematicBreak()
}
