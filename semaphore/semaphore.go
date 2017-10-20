package semaphore

import "sync"

type Semaphore struct {
	Threads chan int
	Wg sync.WaitGroup
}

func NewSemaphore(n int) *Semaphore {
	inst := new(Semaphore)
	inst.Threads = make(chan int,n)
	// inst.Wg=sync.WaitGroup{}
	return inst
}

func (self *Semaphore) P() {
	self.Threads <- 1
	self.Wg.Add(1)
}

func (self *Semaphore) V() {
	self.Wg.Done()
	<-self.Threads
}

func (self *Semaphore) Wait() {
	self.Wg.Wait()
}