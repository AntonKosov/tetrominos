package ticker

import (
	"sync"
	"time"
)

type ticker struct {
	duration     time.Duration
	lastActionAt time.Time
	pausedAt     *time.Time
	doneSignal   chan struct{}
	resume       chan struct{}
	pause        chan struct{}
	reset        chan time.Duration
	waitGroup    sync.WaitGroup
	Tick         chan struct{}
}

func newTicker(duration time.Duration) *ticker {
	t := ticker{
		duration:   duration,
		doneSignal: make(chan struct{}),
		resume:     make(chan struct{}),
		pause:      make(chan struct{}),
		reset:      make(chan time.Duration),
		Tick:       make(chan struct{}),
	}

	t.run()

	return &t
}

func (t *ticker) Done() {
	close(t.doneSignal)
	t.waitGroup.Wait()
	close(t.resume)
	close(t.pause)
	close(t.reset)
	close(t.Tick)
}

func (t *ticker) Pause() {
	t.pause <- struct{}{}
}

func (t *ticker) Resume() {
	t.resume <- struct{}{}
}

func (t *ticker) Reset(d time.Duration) {
	t.reset <- d
}

func (t *ticker) run() {
	t.waitGroup.Add(1)
	go func() {
		defer t.waitGroup.Done()
		t.lastActionAt = *getUTCTime()
		ticker := time.NewTicker(t.duration)
	OuterLoop:
		for {
			select {
			case <-t.doneSignal:
				break OuterLoop
			case <-t.pause:
				t.pausedAt = getUTCTime()
				ticker.Stop()
			case <-t.resume:
				restDuration := t.duration - t.pausedAt.Sub(t.lastActionAt)
				d := t.duration
				if restDuration > 0 {
					d = restDuration
				}
				ticker = time.NewTicker(d)
			case d := <-t.reset:
				t.duration = d
				ticker.Reset(d)
			case <-ticker.C:
				t.lastActionAt = *getUTCTime()
				if t.pausedAt != nil {
					t.pausedAt = nil
					ticker.Reset(t.duration)
				}
				t.Tick <- struct{}{}
			}
		}
		if ticker != nil {
			ticker.Stop()
		}
	}()
}

func getUTCTime() *time.Time {
	t := time.Now().UTC()
	return &t
}
