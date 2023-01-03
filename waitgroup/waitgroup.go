package waitgroup

import "sync"

// A WaitGroup waits for a collection of goroutines to finish.
type WaitGroup struct {
	wg sync.WaitGroup

	pc PanicCatcher
}

// Go to run a function by a goroutine.
func (w *WaitGroup) Go(f func()) {
	w.wg.Add(1)

	go func() {
		defer w.wg.Done()

		w.pc.Try(f)
	}()

}

// Wait for all goroutines to finish.
func (w *WaitGroup) Wait() {
	w.wg.Wait()

	// if some panic happened, we will panic it out.
	w.pc.Repanic()
}
