package document

import (
	"strings"

	"github.com/raykov/mdtopdf/color"
)

// NewSetFontStyleOption sets new Font Style
func (d *Document) NewSetFontStyleOption(fontStyle string) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.FontStyle = fontStyle

		return nil
	}
}

// NewAddFontStyleOption adds Font style to the current
func (d *Document) NewAddFontStyleOption(fontStyle string) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		if strings.Contains(s.FontStyle, fontStyle) {
			return nil
		}
		s.FontStyle += fontStyle

		return nil
	}
}

// NewFontFamilyOption sets Font Family
func (d *Document) NewFontFamilyOption(family string) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.FontFamily = family

		return nil
	}
}

// NewFontSizeOption sets Font Size
func (d *Document) NewFontSizeOption(size float64) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.FontSize = size

		return nil
	}
}

// NewTextColorOption sets Text Color
func (d *Document) NewTextColorOption(color *color.Color) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.TextColor = color

		return nil
	}
}

// NewFillColorOption sets Fill Color
func (d *Document) NewFillColorOption(color *color.Color) func(style any) error {
	return func(style any) error {
		style.(*Style).FillColor = color

		return nil
	}
}

// NewAddMarginOption adds Left Margin
func (d *Document) NewAddMarginOption(margin float64) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.LeftMargin += margin

		return nil
	}
}

// NewPaddingOption Adds Left Margin
func (d *Document) NewPaddingOption(padding float64) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.LeftMargin = padding

		return nil
	}
}

// NewAddCellMarginOption adds CellMargin
func (d *Document) NewAddCellMarginOption(margin float64) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.CellMargin += margin

		return nil
	}
}

// NewSetCellMarginOption sets CellMargin
func (d *Document) NewSetCellMarginOption(margin float64) func(style any) error {
	return func(style any) error {
		s := style.(*Style)
		s.CellMargin = margin

		return nil
	}
}
