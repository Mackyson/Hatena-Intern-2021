package renderer

import (
	"bytes"
	"context"
	"log"

	pb "github.com/hatena/Hatena-Intern-2021/services/renderer-go/pb/fetcher"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"google.golang.org/grpc"
)

type autoTitleLinker struct {
	fetcherCli pb.FetcherClient
	ctx        context.Context
}

func newGoldmarkMDConverter(ctx context.Context) (goldmark.Markdown, error) {
	linker, err := newAutoTitleLinker(ctx)
	if err != nil {
		return nil, err
	}
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(linker, 999),
			),
		),
	)
	return markdown, nil
}

const Address = "localhost:50051" //TODO このマジックナンバーはどうにかなるはず

func newAutoTitleLinker(ctx context.Context) (*autoTitleLinker, error) {
	l := new(autoTitleLinker)
	conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	l.fetcherCli = pb.NewFetcherClient(conn)
	l.ctx = ctx
	return l, nil
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			// title, err := l.fetchTitle(l.ctx, string(node.Destination))
			title, _ := l.fetchTitle(l.ctx, string(node.Destination)) //FIXME errorを握り潰さないと表示してくれない
			// if err != nil {
			// 	return ast.WalkContinue, err
			// // return ast.WalkStop, err
			// }
			node.AppendChild(node, ast.NewString([]byte(title)))
		}
		return ast.WalkContinue, nil
	})
}

func (l *autoTitleLinker) fetchTitle(ctx context.Context, url string) (string, error) {
	//DEBUG
	// return "example", nil
	reply, err := l.fetcherCli.Fetch(ctx, &pb.FetchRequest{Url: url})

	log.Println(reply.GetTitle())
	if err != nil {
		panic(err)
		// return "", err
	}
	//DEBUG
	log.Println(reply.GetTitle())
	return reply.GetTitle(), nil
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var html bytes.Buffer
	markdown, err := newGoldmarkMDConverter(ctx)
	if err != nil {
		return "", err
	}
	err = markdown.Convert([]byte(src), &html)
	return html.String(), err
}
