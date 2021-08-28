package mdtopdf

import (
	"io"

	"github.com/yuin/goldmark/ast"
)

// Writer describes what kind of behaviour we expect from writer
type Writer interface {
	Write(p []byte) (n int, err error)
	Styled
	WriteImage(string, io.Reader, ...ast.Attribute)
	SetLink(string)
	WriteTable(rows ...[]string)
	WithBreaks
}

// Styled describes what kind of behaviour we expect from styled elements
type Styled interface {
	WithStyle(*Style) Writer
	ResetStyle() Writer
	CurrentStyle() *Style
}

// WithBreaks describes what kind of behaviour we expect from breaks
type WithBreaks interface {
	LineBreak()
	ThematicBreak()
}
