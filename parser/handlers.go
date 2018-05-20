package parser

import (
	"sync"
)

type handleVal struct {
	val interface{}
}

var handleLock sync.Mutex
var handleVals = make(map[uintptr]handleVal)
var handleIndex uintptr = 100

func newHandle(v interface{}) uintptr {
	handleLock.Lock()
	defer handleLock.Unlock()
	i := handleIndex
	handleIndex++
	handleVals[i] = handleVal{v}
	return i
}

func lookupHandle(handle uintptr) interface{} {
	handleLock.Lock()
	defer handleLock.Unlock()
	r, ok := handleVals[handle]
	if !ok {
		if handle >= 100 && handle < handleIndex {
			panic("deleted handle")
		} else {
			panic("invalid handle")
		}
	}
	return r.val
}

func deleteHandles() {
	handleLock.Lock()
	defer handleLock.Unlock()
	for handle, _ := range handleVals {
		delete(handleVals, handle)
	}
}
