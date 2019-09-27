package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/kinwyb/blog/application/web/controllers/blog:Blog"] = append(beego.GlobalControllerRouter["github.com/kinwyb/blog/application/web/controllers/blog:Blog"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
