package customer

import (
	"time"

	"github.com/kinwyb/blog/blogModels/blogBeans/dbBeans"
)

//博客列表请求参数
type BlogListReq struct {
	Tp     int64  `description:"类型"`
	IsOpen string `description:"是否开放"`
	Status int64  `description:"状态"`
	Search string `description:"搜索参数"`
}

type BlogListResp struct {
	*dbBeans.Blog
	*dbBeans.BlogStatistics
	TimeAdd time.Time `description:"时间"`
	Tags    []string  `description:"标签"`
}
