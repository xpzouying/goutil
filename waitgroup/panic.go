package waitgroup

import (
	"runtime"
	"runtime/debug"
	"sync/atomic"
)

type PanicCatcher struct {
	recovered atomic.Pointer[RecoveredPanic]
}

func (p *PanicCatcher) Try(f func()) {

	defer p.tryRecover()

	f()
}

func (p *PanicCatcher) tryRecover() {
	if r := recover(); r != nil {
		rp := NewRecoveredPanic(r)

		p.recovered.CompareAndSwap(nil, rp)
	}
}

func (p *PanicCatcher) Repanic() {

	if val := p.recovered.Load(); val != nil {
		panic(val)
	}
}

type RecoveredPanic struct {
	// Value of recover.
	Value any

	Callers []uintptr

	DebugStack []byte
}

func NewRecoveredPanic(r any) *RecoveredPanic {

	var callers [64]uintptr

	n := runtime.Callers(2, callers[:])

	return &RecoveredPanic{
		Value:      r,
		Callers:    callers[:n],
		DebugStack: debug.Stack(),
	}
}
