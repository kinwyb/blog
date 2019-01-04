package main

import (
	_ "github.com/kinwyb/blog/application/web/routers"
	"github.com/kinwyb/blog/blogModels/blogBeans"
	"path/filepath"

	"github.com/kinwyb/blog/blogModels/blogConfig"
	lg "github.com/kinwyb/go/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = filepath.Join(beego.AppPath, "swagger")
	//beego.BeeLogger.SetLevel(beego.LevelWarning)
	if beego.BConfig.RunMode != "dev" {
		logpath := beego.AppConfig.DefaultString("log.path", "")
		//设置日志
		lg.SetLogPath(logpath)
		beego.BeeLogger.Reset()
		beego.BConfig.Log.AccessLogs = true
		//beego.BConfig.Log.AccessLogsFormat = "JSON_FORMAT"
		beego.BeeLogger.Async(1000)
		beego.BeeLogger.SetLogger(logs.AdapterFile,
			"{\"filename\":\""+logpath+"/"+blogBeans.RequestLog+"\",\"level\":7,\"maxlines\":0,\"maxsize\":0,\"daily\":true,\"maxdays\":10}")
		beego.BeeLogger.EnableFuncCallDepth(true)
	}
	blogConfig.InitConfig(beego.AppConfig)
	beego.Run()
}
