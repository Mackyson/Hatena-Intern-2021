package grpc

import (
	"context"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2021/services/renderer-go/pb/renderer"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Render(t *testing.T) {
	s := NewServer()

	tests := []struct {
		src    string
		expect string
	}{
		{
			src: `# hoge
## fuga`,
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
		reply, err := s.Render(context.Background(), &pb.RenderRequest{Src: test.src})
		assert.NoError(t, err)
		assert.Equal(t, test.expect, reply.Html)
	}
}
