package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/raykov/oksvg"
	"github.com/srwiley/rasterx"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

const (
	inlineDataPrefix           = "data:image/"
	incorrectInlineImageFormat = `Incorrect inline image format:`
	missingInlineImageData     = `Missing inline image data:`
	inlineImageError           = `%s
	expected:	![](data:image/png;base64,iVBORw0KG...SUVORK5CYII=)
	got:		![](%s)`
)

func renderImage(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Image)

	filePath := string(util.URLEscape(n.Destination, true))

	switch {
	case strings.HasPrefix(filePath, inlineDataPrefix):
		return inlineData(w, node, filePath)
	case strings.HasPrefix(filePath, "http"):
		return webImage(w, node, filePath)
	default:
		return localImage(w, node, filePath)
	}
}

func inlineData(w Writer, node ast.Node, filePath string) (ast.WalkStatus, error) {
	parts := strings.Split(filePath, ",")
	if len(parts) != 2 {
		log.Printf(inlineImageError, incorrectInlineImageFormat, filePath)
		return ast.WalkSkipChildren, nil
	}

	if parts[1] == "" {
		log.Printf(inlineImageError, missingInlineImageData, filePath)
		return ast.WalkSkipChildren, nil
	}

	partsInfo := strings.Split(strings.TrimPrefix(parts[0], inlineDataPrefix), ";")
	if len(partsInfo) != 2 {
		return ast.WalkSkipChildren, nil
	}

	if partsInfo[1] != "base64" {
		return ast.WalkSkipChildren, nil
	}

	switch partsInfo[0] {
	case "png", "jpeg", "jpg", "gif":
		return inlineImage(w, node, parts[1], partsInfo[0])
	case "svg+xml", "svg":
		return inlineSVG(w, node, parts[1])
	default:
		return ast.WalkSkipChildren, nil
	}
}

func inlineImage(w Writer, node ast.Node, rawData, ext string) (ast.WalkStatus, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return ast.WalkSkipChildren, nil
	}

	file := bytes.NewReader(data)
	w.WriteImage(rawData+"."+ext, file, node.Attributes()...)

	return ast.WalkSkipChildren, nil
}

func inlineSVG(w Writer, node ast.Node, rawData string) (ast.WalkStatus, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return ast.WalkSkipChildren, nil
	}

	icon, err := oksvg.ReadIconStream(bytes.NewReader(data), oksvg.WarnErrorMode)
	if err != nil {
		fmt.Println(err)
		return ast.WalkSkipChildren, nil
	}
	wi, hi := int(icon.ViewBox.W), int(icon.ViewBox.H)
	img := image.NewRGBA(image.Rect(0, 0, wi, hi))

	scannerGV := rasterx.NewScannerGV(wi, hi, img, img.Bounds())
	raster := rasterx.NewDasher(wi, hi, scannerGV)
	icon.Draw(raster, 1.0)
	icon.DrawTexts(img, 1.0)

	var buff = bytes.NewBuffer([]byte{})
	png.Encode(buff, img)

	w.WriteImage(rawData+".png", buff, node.Attributes()...)

	return ast.WalkSkipChildren, nil
}

func webImage(w Writer, node ast.Node, filePath string) (ast.WalkStatus, error) {
	resp, err := http.Get(filePath)
	if err != nil {
		fmt.Printf("File NOT found: %s\n", filePath)
		return ast.WalkSkipChildren, nil
	}
	defer resp.Body.Close()

	w.WriteImage(filePath, resp.Body, node.Attributes()...)

	return ast.WalkSkipChildren, nil
}

func localImage(w Writer, node ast.Node, filePath string) (ast.WalkStatus, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("File NOT found: %s\n", filePath)
		return ast.WalkSkipChildren, nil
	}
	defer file.Close()
	w.WriteImage(filePath, file, node.Attributes()...)

	return ast.WalkSkipChildren, nil
}
