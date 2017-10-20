package concurrent

func (self *ConCurrent) AddError(e error) {
	self.errors = append(self.errors,e)
}

func (self *ConCurrent) GetErrors() []error {
	return self.errors
}

func (self *ConCurrent) GetLastError() error {
	if self.HasError() {
		return self.errors[len(self.errors)-1]
	}
	return nil
}

func (self *ConCurrent) ClearErrors() {
	self.errors = nil
}

func (self *ConCurrent) HasError() bool {
	return len(self.errors) != 0
}