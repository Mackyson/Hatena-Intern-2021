package renderer

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var html bytes.Buffer
	err := goldmark.Convert([]byte(src), &html)
	if err != nil {
		panic(err) //TODO ハンドリング
	}
	return html.String(), nil
}
