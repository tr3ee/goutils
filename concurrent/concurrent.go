package concurrent

import (
	"fmt"
	"reflect"
	"sync"
	"tr3e/utils/semaphore"
)

// Concurrent represents a concurrent operation unit
type Concurrent struct {
	finTicket   uint64
	totalTicket uint64
	orderSignal bool
	closeSignal bool
	cond        *sync.Cond
	semaphore   *semaphore.Semaphore
	errors      []error
	values      chan []interface{}
}

// New returns a new ConCurrent struct pointer
func New(threads int, orderReturn bool) *Concurrent {
	return &Concurrent{
		closeSignal: false,
		orderSignal: orderReturn,
		cond:        sync.NewCond(new(sync.Mutex)),
		semaphore:   semaphore.NewSemaphore(threads),
		values:      make(chan []interface{}, threads),
	}
}

// Call will execute the function concurrently, if the function is executable
func (c *Concurrent) Call(function interface{}, params ...interface{}) error {

	if c.HasError() {
		return c.GetLastError()
	}

	if c.closeSignal {
		return ErrVChanClosed
	}

	f := reflect.TypeOf(function)

	// fucntion callable validate
	if f.Kind() != reflect.Func {
		return fmt.Errorf("Concurrent-handler must be a callable func")
	}

	fin := f.NumIn()
	fout := f.NumOut()

	// params validate
	if fin > len(params) {
		return fmt.Errorf("Call function <%s> with too few input arguments(%d), need %d", function, len(params), fin)
	}
	in := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		in[i] = reflect.ValueOf(params[i])
	}

	// allocate thread
	c.semaphore.P()
	go func(ticket uint64) {
		defer c.semaphore.V() // free thread
		vals := reflect.ValueOf(function).Call(in)
		if fout != len(vals) {
			c.AddError(fmt.Errorf("The number of return values does not match"))
			return
		}
		rets := make([]interface{}, fout)
		for i, v := range vals {
			rets[i] = v.Interface()
		}
		if !c.orderSignal {
			c.values <- rets
			return
		}
		c.cond.L.Lock()
		defer func() {
			c.cond.L.Unlock()
			c.cond.Broadcast()
		}()
		for ticket != c.finTicket {
			c.cond.Wait()
		}
		c.values <- rets
		c.finTicket++
	}(c.totalTicket)
	c.totalTicket++
	return nil
}

// Wait will block until all processes are completed
func (c *Concurrent) Wait() {
	c.semaphore.Wait()
}

// Clear flush the return-value channel
func (c *Concurrent) Clear() {
	for {
		select {
		case <-c.values:
		default:
			return
		}
	}
}

// Close closes the concurrent instance, it will block until all the process are finished
func (c *Concurrent) Close() {
	c.Wait()
	c.closeSignal = true
	close(c.values)
}

// Next will return the value performed by the function, and if it's not prepared yet, Next will block the process
func (c *Concurrent) Next() ([]interface{}, error) {

	value, ok := <-c.values
	if ok {
		return value, nil
	}
	return nil, ErrVChanClosed

}

// NextNoBlock will the value performed by the function without blocking
func (c *Concurrent) NextNoBlock() ([]interface{}, error) {
	select {
	case value, ok := <-c.values:
		if ok {
			return value, nil
		}
		return nil, ErrVChanClosed
	default:
		return nil, ErrNotPrep
	}
}

// isExecutable returns true if its type is reflect.Func
func (c *Concurrent) isExecutable(e interface{}) bool {
	return reflect.TypeOf(e).Kind() == reflect.Func
}
