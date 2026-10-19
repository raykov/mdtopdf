package mdtopdf

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"

	"github.com/raykov/gofpdf"
	"github.com/raykov/mdtopdf/document"
)

type starImager struct{ img *document.InlineImage }

// Segments draws every "★" as an image.
func (s starImager) Segments(text string) []document.Segment {
	var out []document.Segment
	for i, part := range strings.Split(text, "★") {
		if i > 0 {
			out = append(out, document.Segment{Image: s.img})
		}
		if part != "" {
			out = append(out, document.Segment{Text: part})
		}
	}
	return out
}

func testImager(t *testing.T) starImager {
	t.Helper()
	im := image.NewRGBA(image.Rect(0, 0, 4, 4))
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			im.Set(x, y, color.RGBA{R: 255, A: 255})
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, im); err != nil {
		t.Fatal(err)
	}
	return starImager{img: &document.InlineImage{Name: "star", Type: "png", Data: buf.Bytes()}}
}

func convert(t *testing.T, md string, opts ...Option) []byte {
	t.Helper()
	var out bytes.Buffer
	if err := ConvertWithOptions(strings.NewReader(md), &out, opts...); err != nil {
		t.Fatalf("convert: %v", err)
	}
	if !bytes.HasPrefix(out.Bytes(), []byte("%PDF-")) {
		t.Fatalf("output is not a PDF")
	}
	return out.Bytes()
}

func imageCount(pdf []byte) int { return bytes.Count(pdf, []byte("/Subtype /Image")) }

func TestInlineImager(t *testing.T) {
	imager := testImager(t)

	tests := []struct {
		name string
		md   string
	}{
		{"paragraph", "before ★ after"},
		{"heading", "# title ★"},
		{"list", "- item ★"},
		{"link", "[go ★](https://example.com)"},
		{"blockquote", "> quote ★"},
		{"table", "| a | b |\n|---|---|\n| ★ | x |\n"},
		{"many wrap to next line", strings.Repeat("★ ", 400)},
		{"many pages", strings.Repeat("line ★\n\n", 200)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdf := convert(t, tt.md, WithInlineImager(imager))
			if got := imageCount(pdf); got != 1 {
				t.Fatalf("want the image embedded once, got %d", got)
			}
		})
	}
}

func TestWithoutInlineImager(t *testing.T) {
	if got := imageCount(convert(t, "before ★ after")); got != 0 {
		t.Fatalf("want no images, got %d", got)
	}
	// Convert keeps working as before
	var out bytes.Buffer
	if err := Convert(strings.NewReader("plain"), &out); err != nil || out.Len() == 0 {
		t.Fatalf("Convert: %v", err)
	}
}

func TestExtensionsAreApplied(t *testing.T) {
	called := false
	convert(t, "x", WithExtensions(func(*gofpdf.Fpdf) { called = true }))
	if !called {
		t.Fatal("extension not called")
	}
}
