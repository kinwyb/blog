package blogService

import (
	"fmt"
	"strings"
	"time"

	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/blog/blogModels/blogBeans/customer"
	"github.com/kinwyb/blog/blogModels/blogDB"
)

//博客列表
func BlogList(req *customer.BlogListReq, ctx *blogBeans.Context) []*customer.BlogListResp {
	defer ctx.Start("sev.BlogList").Finish()
	result := blogDB.BlogList(req, ctx.Page, ctx.Child())
	var ret []*customer.BlogListResp
	retMap := map[int64]*customer.BlogListResp{}
	var ids []string
	for _, v := range result {
		r := &customer.BlogListResp{
			Blog: v,
		}
		if v.Tag != "" {
			r.Tags = strings.Split(v.Tag, ",")
		}
		r.TimeAdd, _ = time.ParseInLocation(blogBeans.DateTimeFormat, v.TimeAdd, time.Local)
		ids = append(ids, fmt.Sprintf("%d", v.BlogId))
		retMap[v.BlogId] = r
		ret = append(ret, r)
	}
	statistics := blogDB.BlogStatistics(ids, ctx.Child())
	for _, v := range statistics {
		blog := retMap[v.BlogId]
		if blog != nil {
			blog.BlogStatistics = v
		}
	}
	return ret
}
