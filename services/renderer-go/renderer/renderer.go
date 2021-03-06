package renderer

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var markdown = goldmark.New(
	goldmark.WithParserOptions(
		parser.WithASTTransformers(
			util.Prioritized(&autoTitleLinker{}, 999),
		),
	),
)

type autoTitleLinker struct {
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	_ = ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			node.AppendChild(node, ast.NewString([]byte(node.Destination)))
		}
		return ast.WalkContinue, nil
	})
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var html bytes.Buffer
	err := markdown.Convert([]byte(src), &html)
	return html.String(), err
}
