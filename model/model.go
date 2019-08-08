package model

type Article struct { // 单个的文章信息
	SourceURL  string // 资源的路径
	ImageURL   string // 插图的路径
	ImageType  string // 图片的类型
	ImageTotal int    // 需要下载的图片总数
	Id         string //  唯一的ID
	IdSort     int    // 文章的顺序
	Content    string // 文章的内容
	Title      string // 文章的标题
}

type BasicInfo struct { // 基础的信息
	SaveRoot   string // 保存的根路径
	BookName   string
	AuthorName string
}
