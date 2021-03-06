package xerrgroup

import (
	"context"
	"sync"
)

// A group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero group is valid and does not cancel on error.
type group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

type OptionFunc func(o *group)

func NewErrGroup(of ...OptionFunc) *group {
	o := new(group)
	for _, optionFunc := range of {
		optionFunc(o)
	}
	return o
}

func SetCancel(f func()) OptionFunc {
	return func(o *group) {
		o.cancel = f
	}
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &group{cancel: cancel}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
