package concurrent

import "errors"

var ErrNotPrep = errors.New("the next Value is not prepared yet")
var ErrVChanClosed = errors.New("the value-channel has been closed")

func (c *Concurrent) AddError(e error) {
	c.errors = append(c.errors, e)
}

func (c *Concurrent) GetErrors() []error {
	return c.errors
}

func (c *Concurrent) GetLastError() error {
	if c.HasError() {
		return c.errors[len(c.errors)-1]
	}
	return nil
}

func (c *Concurrent) ClearErrors() {
	c.errors = nil
}

func (c *Concurrent) HasError() bool {
	return len(c.errors) != 0
}
