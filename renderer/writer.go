package renderer

import (
	"io"

	"github.com/raykov/mdtopdf/color"

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
	NewSetFontStyleOption(string) func(style interface{}) error
	NewAddFontStyleOption(string) func(style interface{}) error
	NewFontFamilyOption(string) func(style interface{}) error
	NewFontSizeOption(float64) func(style interface{}) error

	NewTextColorOption(*color.Color) func(style interface{}) error
	NewFillColorOption(*color.Color) func(style interface{}) error

	NewAddMarginOption(float64) func(style interface{}) error
	NewAddCellMarginOption(float64) func(style interface{}) error
	NewSetCellMarginOption(float64) func(style interface{}) error
	NewPaddingOption(float64) func(style interface{}) error

	WithStyle(...func(style interface{}) error)
	ResetStyle()
}

// WithBreaks describes what kind of behaviour we expect from breaks
type WithBreaks interface {
	LineBreak()
	ThematicBreak()
}
