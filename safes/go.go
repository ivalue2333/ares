package safes

import (
	"runtime/debug"

	"github.com/ivalue2333/ares/logs"
)

func Go(do func(), recoverDo func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				recoverDo(err)
				return
			}
		}()
		do()
	}()
}

func GoLog(do func()) {
	Go(do, func(err interface{}) {
		logs.Info("SafeGoFmt panic, err:%+v, stack:\n%s", err, string(debug.Stack()))
	})
}
