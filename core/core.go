package core

import (
	"fmt"
	"get-data-from-web-make-ebook/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var contentUrl = "https://www.wenku8.net/novel/0/566/18663.htm"

func TitleL1(title string) string { return "# " + title }
func TitleL2(title string) string { return "## " + title }
func TitleL3(title string) string { return "### " + title }
func Image(id string) string      { return fmt.Sprintf("![](%s)\n", id) }

// 获取文章的内容，并去掉广告
func GetContentFromInput(input io.Reader) string {
	reader := transform.NewReader(input, simplifiedchinese.GBK.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	contentDoc := doc.Find("#content")
	contentDoc.Children().RemoveFiltered("#contentdp")
	return contentDoc.Text()
}

// 解析标题，变成markdown形式的标题以及图片链接
func ParseContent(article model.Article) model.Article {

	strs := strings.Split(article.Content, "\n")
	imageCount := 0
	for i, v := range strs {
		strs[i] = (func(line string) string {
			line = strings.TrimSpace(line)

			// 此为二级标题，例： 第一章 ... ##
			if matched, _ := regexp.Match("^第.*章", []byte(line)); matched {
				return TitleL2(line)
			}

			// 此为三级标题，例：断章...  ###
			if matched, _ := regexp.Match("^断章", []byte(line)); matched {
				return TitleL3(line)
			}

			// 此为图片，例：插图 ...
			if matched, _ := regexp.Match("^插图", []byte(line)); matched {
				imageCount++
				return Image(path.Join(article.Id, strconv.Itoa(imageCount)+"."+article.ImageType))
			}
			return line
		})(v)
	}

	article.Content = strings.Join(strs, "\n")
	article.ImageTotal = imageCount
	return article
}

// 下载保存图片
func FetchAndSaveImage(article model.Article, basicInfo model.BasicInfo) {
	log.Printf("Fetching image for article %s\n", article.Title)

	// get page from web
	res, err := http.Get(article.ImageURL)
	if err != nil {
		log.Fatalf("error when fetch image page. %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// parse html page
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// get image url
	var imageSrcs []string
	for _, node := range doc.Find(".divimage a img").Nodes {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				imageSrcs = append(imageSrcs, attr.Val)
				break
			}
		}
	}
	imageSrcs = imageSrcs[len(imageSrcs)-article.ImageTotal:]

	// download pages
	var wg sync.WaitGroup
	wg.Add(len(imageSrcs))

	for imgID, src := range imageSrcs {
		// 限速
		<-time.Tick(time.Second)
		go func(wg *sync.WaitGroup, src string, imgID int) {
			log.Printf("Fetching image from %s\n", src)

			if res, err := http.Get(src); err != nil {
				log.Printf("error when fetch image %s with article id %s\n", src, article.Id)
			} else {
				defer func() {
					_ = res.Body.Close()
					wg.Done()
				}()
				// make directory
				_ = os.Mkdir(path.Join(basicInfo.SaveRoot, article.Id), 0777)

				// save image
				imgPath := path.Join(basicInfo.SaveRoot, article.Id, strconv.Itoa(imgID)+"."+article.ImageType)
				if data, err := ioutil.ReadAll(res.Body); err != nil {
					log.Printf("error when save image %s with article id %s\n", src, article.Id)
				} else {
					_ = ioutil.WriteFile(imgPath, data, 0777)
				}
			}
		}(&wg, src, imgID+1)
	}
	wg.Wait()
}

func GenerateEpub(articleInfos []model.Article, basicInfo model.BasicInfo) {
	tmpFileName := path.Join(basicInfo.SaveRoot, basicInfo.BookName+".txt")
	f, err := os.Create(tmpFileName)
	if err != nil {
		log.Fatalf("create file error. %v", err)
	}
	defer f.Close()
	_, _ = f.WriteString("% " + basicInfo.BookName + "\n")
	_, _ = f.WriteString("% " + basicInfo.AuthorName + "\n")

	for _, article := range articleInfos {
		_, _ = f.WriteString(TitleL1(article.Title) + "\n")
		_, _ = f.WriteString(article.Content + "\n")
	}

	cmd := exec.Command("pandoc", tmpFileName, "-o", path.Join(basicInfo.SaveRoot, basicInfo.BookName+".epub"))
	cmd.Dir = basicInfo.SaveRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		log.Fatalf("Error when generate epub file. %v", err)
	}
}
