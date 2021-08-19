package fetcher

import (
	"context"
	"log"
)

func Fetch(ctx context.Context, url string) (string, error) {
	log.Println("passed")
	return "tmp", nil //TODO URLをたどってtitleタグの中身を抜き出す部分の実装
}
