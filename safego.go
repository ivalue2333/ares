package ares

import (
	"runtime/debug"
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
		Info("SafeGoFmt panic, err:%+v, stack:\n%s", err, string(debug.Stack()))
	})
}
