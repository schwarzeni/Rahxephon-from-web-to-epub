package main

import (
	"get-data-from-web-make-ebook/core"
	"get-data-from-web-make-ebook/model"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(len(articleInfos))

	// 读取配置文件
	for idx, article := range articleInfos {
		go func(article model.Article, articleIdx int, wg *sync.WaitGroup) {
			log.Printf("Working on %s from %s\n", article.Title, article.SourceURL)
			res, err := http.Get(article.SourceURL)
			if err != nil {
				log.Fatalf("error when fetch article page. %v", err)
			}
			defer func() {
				wg.Done()
				res.Body.Close()
			}()

			article.Content = core.GetContentFromInput(res.Body)
			article = core.ParseContent(article)
			if article.ImageTotal > 0 {
				core.FetchAndSaveImage(article, basicInfo)
			}
			articleInfos[articleIdx] = article
			log.Printf("Finish working on %s from %s\n", article.Title, article.SourceURL)
		}(article, idx, &wg)
	}

	wg.Wait()
	log.Printf("All is finish! Concating files\n")

	// 生成epub文件
	core.GenerateEpub(articleInfos, basicInfo)

}
