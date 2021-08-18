package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
	tests := []struct {
		src    string
		expect string
	}{
		{
			src: "# hoge\n## fuga",
			expect: `<h1>hoge</h1>
<h2>fuga</h2>
`,
		},
		{
			src: `* a
* b
1. 1
1. 2`,
			expect: `<ul>
<li>a</li>
<li>b</li>
</ul>
<ol>
<li>1</li>
<li>2</li>
</ol>
`,
		},
		{
			src: `[google](https://google.com/)`,
			expect: `<p><a href="https://google.com/">google</a></p>
`,
		},
	}
	for _, test := range tests {
		html, err := Render(context.Background(), test.src)
		assert.NoError(t, err)
		assert.Equal(t, test.expect, html)
	}
}
