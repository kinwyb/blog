package blogBeans

import (
	"github.com/kinwyb/go/logs"
)

const ServiceLog = "service.log"
const RequestLog = "request.log"
const DBLog = "db.log"

var log = logs.GetLogger(ServiceLog)

func init() {
	logs.RegisterLog(setLog)
}

func setLog(lg *logs.LogFiles) {
	log = lg.GetLog(ServiceLog)
}
