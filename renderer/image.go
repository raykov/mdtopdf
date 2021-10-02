package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

const inlineImg = "data:image/"

func renderImage(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Image)

	filePath := string(util.URLEscape(n.Destination, true))

	switch {
	case strings.HasPrefix(filePath, inlineImg):
		return inlineImage(w, node, filePath)
	case strings.HasPrefix(filePath, "http"):
		return webImage(w, node, filePath)
	default:
		return localImage(w, node, filePath)
	}
}

func inlineImage(w Writer, node ast.Node, filePath string) (ast.WalkStatus, error) {
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
