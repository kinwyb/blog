package blogService

import (
	"github.com/kinwyb/blog/blogModels/blogBeans"
	"github.com/kinwyb/go/logs"
)

var log = logs.GetLogger(blogBeans.ServiceLog)

func init() {
	logs.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(blogBeans.ServiceLog)
}
