package semaphore

import (
	"testing"
	"math/rand"
	"fmt"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(3)

	for{
		s.P()
		go func(){
			t := rand.Intn(5)
			fmt.Println("sub called,sleeping",t)
			time.Sleep(time.Duration(t)*time.Second)
			s.V()
		}()
	}

}