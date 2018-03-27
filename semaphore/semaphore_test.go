package semaphore

import (
	"math/rand"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(5)

	for index := 0; index < 20; index++ {
		s.P()
		go func() {
			r := rand.Intn(3)
			t.Log("sub called,sleeping", r)
			time.Sleep(time.Duration(r) * time.Second)
			s.V()
		}()
	}
}
