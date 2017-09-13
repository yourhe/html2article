package main

import "github.com/yourhe/html2article"

func main() {
	urlStr := "http://www.fao.org/in-action/blazing-a-trail-for-timber-traceability-in-benin/en/"
	ext, err := html2article.NewFromUrl(urlStr)
	if err != nil {
		panic(err)
	}
	article, err := ext.ToArticle()
	if err != nil {
		panic(err)
	}
	println("article title is =>", article.Title)
	println("article publishtime is =>", article.Publishtime) //using UTC timezone
	println("article content is =>", article.Content)
	println("article images is =>", article.Images)

	//parse the article to be readability
	article.Readable(urlStr)
	println("read=>", article.ReadContent)
}
