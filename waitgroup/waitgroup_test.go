package waitgroup

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleWaitGroup() {
	var count atomic.Uint64

	var wg WaitGroup
	for i := 0; i < 100; i++ {

		wg.Go(func() {
			count.Add(1)
		})
	}
	wg.Wait()

	fmt.Println(count.Load())
	// Output:
	// 100
}

func TestWaitGroup(t *testing.T) {

	var count atomic.Uint64

	var wg WaitGroup
	for i := 0; i < 100; i++ {

		wg.Go(func() {
			count.Add(1)
		})
	}
	wg.Wait()

	want := uint64(100)
	got := count.Load()
	require.Equal(t, want, got)
}

func TestWaitGroupPanic(t *testing.T) {

	t.Run("run with panic function", func(t *testing.T) {

		var wg WaitGroup

		wg.Go(func() {
			panic("some panic happened")
		})

		require.Panics(t, wg.Wait)
	})

}
