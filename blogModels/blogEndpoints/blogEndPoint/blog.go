package blogEndPoint

import (
	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/blog/blogModels/blogBeans/customer"
	"github.com/kinwyb/blog/blogModels/blogEndpoints"
	"github.com/kinwyb/blog/blogModels/blogService"
	"github.com/kinwyb/go/err1"
)

//blog首页
func Index(search string, ctx *blogBeans.Context) err1.Error {
	defer ctx.Start("ep.Blog.Index")
	if err := blogEndpoints.CheckPower("", ctx.Child()); err != nil {
		return err
	}
	req := &customer.BlogListReq{
		Search: search,
		Tp:     0,
		IsOpen: "1",
		Status: 99,
	}
	ctx.Data["data"] = blogService.BlogList(req, ctx.Child())
	ctx.Template = "blog/index.html"
	return nil
}
