package blogBeans

type BlogType int64

const (
	BlogTypeArticle BlogType = iota //文章
)

func (b BlogType) Int64() int64 {
	return int64(b)
}

func (b BlogType) String() string {
	switch b {
	case BlogTypeArticle:
		return "文章"
	default:
		return ""
	}
}

type BlogStatus int64

const (
	BlogStatusRelease BlogStatus = iota + 1 //博客状态
)

func (b BlogStatus) Int64() int64 {
	return int64(b)
}

func (b BlogStatus) String() string {
	switch b {
	case BlogStatusRelease:
		return "发布"
	default:
		return ""
	}
}
