package concurrent

import (
	"fmt"
	"reflect"
	"sync"
	"tr3e/utils/semaphore"
)

type Executable interface{}

type ReturnType []interface{}

type ConCurrent struct {
	rcnt            uint64
	curr            uint64
	nextIterator    uint64
	nextBlockSignal bool
	nextChannel     chan int
	nextMutex       *sync.Mutex
	closeSignal     bool
	cond            *sync.Cond
	semaphore       *semaphore.Semaphore
	errors          []error
	Values          []([]interface{})
}

func New(threads int) *ConCurrent {
	return &ConCurrent{
		closeSignal:     false,
		nextBlockSignal: false,
		nextMutex:       new(sync.Mutex),
		nextChannel:     make(chan int, 1),
		cond:            sync.NewCond(new(sync.Mutex)),
		semaphore:       semaphore.NewSemaphore(threads),
	}
}

func (self *ConCurrent) isExecutable(e Executable) bool {
	return reflect.TypeOf(e).Kind() == reflect.Func
}

func (self *ConCurrent) Call(function Executable, params ...interface{}) error {

	if self.HasError() {
		return self.GetLastError()
	}

	if self.closeSignal {
		return fmt.Errorf("ConCurrent: Channel have been closed")
	}

	f := reflect.TypeOf(function)

	// fucntion callable validate
	if f.Kind() != reflect.Func {
		return fmt.Errorf("ConCurrent: handler must be a callable func")
	}

	fin := f.NumIn()
	fout := f.NumOut()

	// params validate
	if fin > len(params) {
		return fmt.Errorf("ConCurrent: Call function <%s> with too few input arguments(%d), need %d", function, len(params), fin)
	}
	in := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		in[i] = reflect.ValueOf(params[i])
	}

	// allocate thread
	self.semaphore.P()
	go func(rId uint64) {
		defer self.semaphore.V() // free thread
		vals := reflect.ValueOf(function).Call(in)
		if fout != len(vals) {
			self.AddError(fmt.Errorf("ConCurrent: The number of return values does not match"))
			return
		}
		rets := make(ReturnType, fout)
		for i, v := range vals {
			rets[i] = v.Interface()
		}
		self.cond.L.Lock()
		defer func() {
			self.cond.L.Unlock()
			self.cond.Broadcast()
		}()
		for self.curr != rId {
			self.cond.Wait()
		}
		self.Values = append(self.Values, rets)
		self.curr++

		self.nextMutex.Lock()
		if self.nextBlockSignal {
			self.nextBlockSignal = false
			self.nextChannel <- 1
		}
		self.nextMutex.Unlock()
	}(self.rcnt)
	self.rcnt++
	return nil
}

func (self *ConCurrent) Wait() {
	self.semaphore.Wait()
}

func (self *ConCurrent) Clear() {
	self.Values = nil
}

func (self *ConCurrent) Next(block bool) ([]interface{}, error) {
	var value []interface{} = nil
	var ok bool = false
	if self.nextIterator < self.curr {
		value = self.Values[self.nextIterator]
		ok = true
	}
	if block {
		if !ok {
			self.nextBlockSignal = true
			<-self.nextChannel
			return self.Next(true)
		}
		self.nextIterator++
		return value, nil
	} else {
		if !ok {
			return nil, fmt.Errorf("ConCurrent: Next Value is not prepared yet.")
		}
		self.nextIterator++
		return value, nil
	}
}
