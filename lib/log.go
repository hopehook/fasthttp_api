package lib

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

var Logger = logs.GetBeeLogger()

func init() {
	conf := fmt.Sprintf(`{
		  "filename": "%s",
	      "maxdays": %s,
	      "daily": %s,
		  "rotate": %s,
	      "level": %s,
	      "separate": %s
	 }`, LOG["path"], LOG["maxdays"], "true", "true", LOG["level"], LOG["separate"])
	logs.SetLogger(logs.AdapterMultiFile, conf)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
}
