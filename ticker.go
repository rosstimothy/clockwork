package clockwork

import (
	"time"
)

// Ticker provides an interface which can be used instead of directly
// using the ticker within the time module. The real-time ticker t
// provides ticks through t.C which becomes now t.Chan() to make
// this channel requirement definable in this interface.
type Ticker interface {
	Chan() <-chan time.Time
	Stop()
}

type realTicker struct{ *time.Ticker }

func (rt realTicker) Chan() <-chan time.Time {
	return rt.C
}

type fakeTicker struct {
	fakeTimer
}

func (ft *fakeTicker) Stop() {
	// Ignore returned bool to make signature match.
	ft.fakeTimer.Stop()
}
