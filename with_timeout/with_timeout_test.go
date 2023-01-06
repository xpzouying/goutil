package with_timeout

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGo(t *testing.T) {

	t.Run("finish without error", func(t *testing.T) {

		var finish atomic.Bool

		f := func() {
			time.Sleep(20 * time.Millisecond)
			finish.Store(true)
		}

		err := Go(f, 100*time.Millisecond)
		assert.NoError(t, err)

		want := true
		got := finish.Load()
		assert.Equal(t, want, got)

	})

	t.Run("finish with timeout error", func(t *testing.T) {

		var finish atomic.Bool

		f := func() {
			time.Sleep(20 * time.Millisecond)
			finish.Store(true)
		}

		err := Go(f, 10*time.Millisecond)
		assert.ErrorIs(t, err, ErrTimeout)

		want := false
		got := finish.Load()
		assert.Equal(t, want, got)

	})
}
