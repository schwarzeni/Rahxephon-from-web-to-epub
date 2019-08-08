package core

import (
	"fmt"
	"get-data-from-web-make-ebook/model"
	"os"
	"testing"
)

const TEST_FILE_PATH = "testwebpage2.html"

// 解析出文章内容，并去除广告
func TestFetchContent(t *testing.T) {
	b, _ := os.Open(TEST_FILE_PATH)

	fmt.Println(GetContentFromInput(b))
}

// 解析标题，变成markdown形式的标题以及图片链接
func TestParseContent(t *testing.T) {
	b, _ := os.Open(TEST_FILE_PATH)
	content := GetContentFromInput(b)

	article := model.Article{
		SourceURL: "",
		ImageURL:  "",
		ImageType: "jpg",
		Id:        "1",
		Content:   content,
	}

	result := ParseContent(article)
	fmt.Println(result)
}

func TestFetchAndSaveImage(t *testing.T) {
	FetchAndSaveImage(model.Article{
		SourceURL:  "",
		ImageURL:   "https://www.wenku8.net/modules/article/reader.php?aid=566&cid=24717",
		ImageType:  "jpg",
		ImageTotal: 6,
		Id:         "1",
		IdSort:     0,
		Content:    "",
	}, model.BasicInfo{SaveRoot: "/Users/nizhenyang/my-project/back-end/golang/get-data-from-web-make-ebook/epub"})
}
