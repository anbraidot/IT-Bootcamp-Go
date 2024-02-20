package positioner

// NewPositionerStub creates a new PositionerStub
func NewPositionerStub() *PositionerStub {
	defaultFuncGetLinearDistance := func(from, to *Position) (linearDistance float64) {
		linearDistance = 0
		return
	}
	return &PositionerStub{
		FuncGetLinearDistance: defaultFuncGetLinearDistance,
	}
}

// PositionerStub is a stub implementation of the Positioner interface
type PositionerStub struct {
	// FuncGetLinearDistance is the function that will be called when GetLinearDistance is called
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = p.FuncGetLinearDistance(from, to)
	return
}