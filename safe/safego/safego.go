package safego

import (
	"runtime/debug"

	"github.com/ivalue2333/ares"
)

func SafeGo(do func(), recoverDo func(err interface{})) {
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

func SafeGoFmt(do func()) {
	SafeGo(do, func(err interface{}) {
		ares.Info("SafeGoFmt panic, err:%+v, stack:\n%s", err, string(debug.Stack()))
	})
}
