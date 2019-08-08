package main

import "get-data-from-web-make-ebook/model"

var basicInfo = model.BasicInfo{
	SaveRoot:   "/Users/nizhenyang/my-project/back-end/golang/get-data-from-web-make-ebook/epub",
	BookName:   "翼神世音",
	AuthorName: "大野木宽",
}

var articleInfos = []model.Article{
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18663.htm",
		ImageURL:   "",
		ImageType:  "jpg",
		ImageTotal: 0,
		Id:         "1",
		IdSort:     1,
		Content:    "",
		Title:      "第一卷",
	},
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18665.htm",
		ImageURL:   "",
		ImageType:  "",
		ImageTotal: 0,
		Id:         "2",
		IdSort:     2,
		Content:    "",
		Title:      "第二卷",
	},
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18667.htm",
		ImageURL:   "https://www.wenku8.net/novel/0/566/24715.htm",
		ImageType:  "jpg",
		ImageTotal: 0,
		Id:         "3",
		IdSort:     3,
		Content:    "",
		Title:      "第三卷",
	},
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18669.htm",
		ImageURL:   "https://www.wenku8.net/novel/0/566/24716.htm",
		ImageType:  "jpg",
		ImageTotal: 0,
		Id:         "4",
		IdSort:     4,
		Content:    "",
		Title:      "第四卷",
	},
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18671.htm",
		ImageURL:   "https://www.wenku8.net/novel/0/566/24717.htm",
		ImageType:  "jpg",
		ImageTotal: 0,
		Id:         "5",
		IdSort:     5,
		Content:    "",
		Title:      "第五卷",
	},
	{
		SourceURL:  "https://www.wenku8.net/novel/0/566/18673.htm",
		ImageURL:   "",
		ImageType:  "jpg",
		ImageTotal: 0,
		Id:         "6",
		IdSort:     6,
		Content:    "",
		Title:      "外传 梦之卵",
	},
}
