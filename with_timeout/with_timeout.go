package with_timeout

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("run in timeout")

func Go(f func(), duration time.Duration) error {

	done := make(chan struct{}, 1)
	go func() {
		defer close(done)

		f()
	}()

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	select {
	case <-ticker.C:
		return ErrTimeout
	case <-done:
		return nil
	}

}
