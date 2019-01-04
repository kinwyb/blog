package blogDB

import (
	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/go/logs"
)

var log = logs.GetLogger(blogBeans.DBLog)

func init() {
	logs.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(blogBeans.DBLog)
}

//获取日志
func GetLog() logs.Logger {
	return log
}