package tgraph

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"io/ioutil"
)

func PublishHtml(sourceTitle string, title string, rawLink string, html string) (string, error) {

     headdata, err := ioutil.ReadFile("/root/flowerss-bot/tgraph/abc.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		headd := string(headdata)
		return headd,err
	}
	headd := string(headdata)

	html = html + fmt.Sprintf(
		"<hr>%s<hr><p>本文章由 <a href=\"https://www.443344.xyz\">Alure资源社</a> 收集整理。</p><p>加入tg频道：<a href=\"https://t.me/best00000\">@best00000</p>",
		headd,
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

