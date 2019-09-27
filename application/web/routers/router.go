// @APIVersion 2.0.0
// @Title blog
// @Description blog
// @Contact wangyb@zhifangw.cn
// @TermsOfServiceUrl http://www.zhifangw.cn
package routers

import (
	"github.com/kinwyb/blog/application/web/controllers/blog"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&blog.Blog{})
}
