package prey

import "06-Testing/02-Double-Test/positioner"

// PreyStub is a stub of the Prey interface
type PreyStub struct {
	// FuncGetSpeed is the function that will be called when GetSpeed is called
	FuncGetSpeed func() (speed float64)
	// FuncGetPosition is the function that will be called when GetPosition is called
	FuncGetPosition func() (position *positioner.Position)
}

// GetSpeed returns the speed of the prey
func (p *PreyStub) GetSpeed() (speed float64) {
	speed = p.FuncGetSpeed()
	return
}

// GetPosition returns the position of the prey
func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = p.FuncGetPosition()
	return
}