package simulator

// NewCatchSimulatorMock creates a new CatchSimulatorMock
func NewCatchSimulatorMock() *CatchSimulatorMock {
	defaultFuncCanCatch := func(hunter, prey *Subject) (canCatch bool) {
		return false
	}

	return &CatchSimulatorMock{
		FuncCanCatch: defaultFuncCanCatch,
	}
}

// Spy is a struct that represents a spy
type Spy struct {
	// NumCalls is the number of times that the function was called
	NumCalls int
	// Args is the list of arguments that the function was called with
	Args []interface{}
}

// CatchSimulatorMock is a mock implementation of CatchSimulator
type CatchSimulatorMock struct {
	// FuncCanCatch is the function that will be called when CanCatch is called
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
	// Spy
	Spy
}

// CanCatch returns true if the hunter can catch the prey
func (c *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	// increment number of calls
	c.NumCalls++
	// append arguments
	c.Args = append(c.Args, []interface{}{hunter, prey})
	// call function
	canCatch = c.FuncCanCatch(hunter, prey)
	return
}
