package renderer

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
)

// Render は受け取った文書を HTML に変換する
unc Render(ctx context.Context, src string) (string, error) {
	var html bytes.Buffer
	err := goldmark.Convert([]byte(src), &html)
	return html.String(), err
}
