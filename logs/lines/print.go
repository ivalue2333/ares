package lines

import "github.com/ivalue2333/ares/logs"

const (
	lineStr = "-------------------------------"
)

func Start() {
	logs.Info(lineStr)
}

func StartEnd() func() {
	logs.Info(lineStr)
	return func() {
		logs.Info(lineStr)
	}
}
