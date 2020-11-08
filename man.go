package manpage

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func writeLines(w util.BufWriter, source []byte, n ast.Node) {
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		_, _ = w.Write(line.Value(source))
	}
}

func (r *man) renderDocument(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)
	l := n.Level
	if l > 2 {
		l = 2
	}
	switch r.Config.Format {
	case SCDOC:
		if entering {
			if n.PreviousSibling() != nil {
				_, _ = w.WriteString("\n\n")
			}
			prefix := strings.Repeat("#", l)
			_, _ = w.WriteString(prefix)
			_ = w.WriteByte(' ')
		} else {
			_ = w.WriteByte('\n')
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderBlockquote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	switch r.Config.Format {
	case SCDOC:
		if entering {
			_ = w.WriteByte('\n')
			_, _ = w.WriteString("```")
			_ = w.WriteByte('\n')
			writeLines(w, source, node)
		} else {
			_, _ = w.WriteString("```")
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return r.renderCodeBlock(w, source, node, entering)
}

func (r *man) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.List)
	if n.IsOrdered() && n.HasChildren() {
		n.FirstChild().SetAttribute([]byte("ordered"), true)
	}
	return ast.WalkContinue, nil
}

func (r *man) renderListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.ListItem)
	var odererd bool
	if _, ok := n.Attribute([]byte("ordered")); ok {
		if n.NextSibling() != nil {
			n.NextSibling().SetAttribute([]byte("ordered"), true)
		}
		odererd = true
	}
	switch r.Config.Format {
	case SCDOC:
		if entering {
			if odererd {
				_, _ = w.WriteString(". ")
			} else {
				_, _ = w.WriteString("- ")
			}
		} else {
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderParagraph(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	switch r.Config.Format {
	case SCDOC:
		if entering {
			_ = w.WriteByte('\n')
		} else {
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderTextBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	switch r.Config.Format {
	case SCDOC:
		if !entering {
			if _, ok := n.NextSibling().(ast.Node); ok && n.FirstChild() != nil {
				_ = w.WriteByte('\n')
			}
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderThematicBreak(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return r.renderEmphasis(w, source, node, entering)
}

func (r *man) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	switch r.Config.Format {
	case SCDOC:
		if entering {
			_ = w.WriteByte('_')
		} else {
			_ = w.WriteByte('_')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	switch r.Config.Format {
	case SCDOC:
		n.RemoveChildren(node)
		if entering {
			_, _ = w.WriteString("<_")
			_, _ = w.Write(n.Destination)
		} else {
			_, _ = w.WriteString("_>")
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *man) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Text)
	s := n.Segment
	if n.IsRaw() {
		w.Write(s.Value(source))
	} else {
		w.Write(s.Value(source))
		if n.HardLineBreak() || n.SoftLineBreak() {
			w.WriteRune('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *man) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Text)
	s := n.Segment
	w.Write(s.Value(source))
	return ast.WalkContinue, nil
}

func renderVerbatim(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if _, err := w.Write(node.Text(source)); err != nil {
			return ast.WalkStop, err
		}
	}
	return ast.WalkContinue, nil
}
