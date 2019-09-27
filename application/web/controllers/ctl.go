package controllers

import (
	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/go"
	"github.com/kinwyb/go/db"
	"github.com/kinwyb/go/db/mysql"

	"time"

	"github.com/astaxie/beego"
	"github.com/rcrowley/go-metrics"
)

//Controller 接口控制器
type TplController struct {
	beego.Controller
	st   time.Time
	OCtx *blogBeans.Context
}

func (ctl *TplController) Prepare() {
	ctl.OCtx = blogBeans.NewContextWithTracing(
		heldiamgo.NewTracingSpanExtractHttpRequest(ctl.Ctx.Request))
	ctl.OCtx.Token = ctl.Ctx.Input.Header("token")
	ctl.OCtx.Query = mysql.GetDBConnect()
	ctl.OCtx.Data = ctl.Controller.Data
	pg := ctl.Ctx.Input.Param(":page")
	if pg != "" {
		p, err := db.Int(pg)
		if err == nil {
			ctl.OCtx.Page = &db.PageObj{
				Page: p,
				Rows: 20,
			}
		}
	}
	ctl.st = time.Now()
}

// Render sends the response with rendered template bytes as text/html type.
func (ctl *TplController) Render() error {
	//性能统计处理
	defer func(ctl *TplController) {
		if enableMetrics {
			metrics.GetOrRegisterTimer(ctl.Ctx.Request.URL.Path, metricsRegistry).UpdateSince(ctl.st)
			//总数请求
			metrics.GetOrRegisterTimer(allRequestMetrics, metricsRegistry).UpdateSince(ctl.st)
		}
	}(ctl)
	//todo: 错误提示错误
	ctl.Data = ctl.OCtx.Data
	ctl.TplName = "blog/" + ctl.OCtx.Template
	if ctl.OCtx.Page != nil {
		ctl.Data["page"] = ctl.OCtx.Page
	}
	return ctl.Controller.Render()
}

//错误日志
func (ctl *TplController) LogError(format string, args ...interface{}) {
	log.Error(format, args...)
}
