package renderer

import "github.com/yuin/goldmark/ast"

func renderTable(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func renderTableHeader(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {

	return ast.WalkContinue, nil
}

func renderTableRow(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func renderTableCell(w Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}
