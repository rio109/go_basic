package bucket

import (
	"sync"
	"time"
	"whatisyourtime/internal/constants"
)

type Bucket struct {
	interval time.Duration // 요청 간격
	lastTime time.Time     // 마지막 요청 시간
	mu       sync.Mutex
}

func New(rate int) *Bucket {
	interval := time.Second / time.Duration(rate)
	return &Bucket{
		interval: interval,
		lastTime: time.Now(),
	}
}

func (b *Bucket) TryAcquire() {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastTime)
	waitTime := b.interval - elapsed

	if waitTime > constants.None {
		time.Sleep(waitTime)
		now = time.Now()
	} else {
		now = b.lastTime.Add(b.interval)
	}
	b.lastTime = now
}
