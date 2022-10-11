package renderer

import (
	"io"

	"github.com/Acke0/mdtopdf/color"

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
	NewSetFontStyleOption(string) func(style any) error
	NewAddFontStyleOption(string) func(style any) error
	NewFontFamilyOption(string) func(style any) error
	NewFontSizeOption(float64) func(style any) error

	NewTextColorOption(*color.Color) func(style any) error
	NewFillColorOption(*color.Color) func(style any) error

	NewAddMarginOption(float64) func(style any) error
	NewAddCellMarginOption(float64) func(style any) error
	NewSetCellMarginOption(float64) func(style any) error
	NewPaddingOption(float64) func(style any) error

	WithStyle(...func(style any) error)
	ResetStyle()
}

// WithBreaks describes what kind of behaviour we expect from breaks
type WithBreaks interface {
	LineBreak()
	ThematicBreak()
}
