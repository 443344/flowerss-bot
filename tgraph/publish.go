package tgraph

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func PublishHtml(sourceTitle string, title string, rawLink string, html string) (string, error) {
	//html = fmt.Sprintf(
	//	"<p>本文章由 <a href=\"https://www.443344.xyz\">Alure资源社</a> 发布。</p><hr>",
	//) + html + fmt.Sprintf(
	//	"<hr><p>本文章由 <a href=\"https://www.443344.xyz\">Alure资源社</a> 发布。</p><p>查看原文：<a href=\"%s\">%s - %s</p>",
	//	rawLink,
	//	title,
	//	sourceTitle,
	//)

	html = html + fmt.Sprintf(
		"<hr><p>本文章由 <a href=\"https://www.443344.xyz\">Alure资源社</a> 发布。</p><p>查看原文：<a href=\"%s\">%s - %s</p>",
		rawLink,
		title,
		sourceTitle,
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	client := clientPool[rand.Intn(len(clientPool))]

	if page, err := client.CreatePageWithHTML(title+" - "+sourceTitle, sourceTitle, rawLink, html, true); err == nil {
		log.Printf("Created telegraph page url: %s", page.URL)
		return page.URL, err
	} else {
		log.Printf("Create telegraph page error: %s", err)
		return "", nil
	}
}
