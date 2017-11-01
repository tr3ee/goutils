package concurrent

import "testing"

import "time"
import "gopkg.in/cheggaaa/pb.v1"

func TestConcurrent(t *testing.T) {
	c := New(200, true)
	retInt := func(i int) int {
		time.Sleep(time.Duration((i%3 + 1)) * time.Second)
		return i
	}
	ThreadCount := 1000
	pushbar := pb.StartNew(ThreadCount)
	pullbar := pb.StartNew(ThreadCount)
	pool, _ := pb.StartPool(pushbar, pullbar)
	go func() {
		for i := 0; i < ThreadCount; i++ {
			if err := c.Call(retInt, i); err != nil {
				panic(err)
			}
			pushbar.Increment()
		}
		pushbar.FinishPrint("Push over")
		c.Close()
	}()
	for {
		_, err := c.Next()
		if err == ErrVChanClosed {
			break
		}
		if err != nil {
			panic(err)
		}
		pullbar.Increment()
	}
	pullbar.FinishPrint("Pull over")
	pool.Stop()
}
