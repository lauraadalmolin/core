package world_model

type Position struct {
	x float64
	y float64
}

// GetX returns the x position
func (position Position) GetX() float64 {
	return position.x
}

// GetY returns the y position
func (position Position) GetY() float64 {
	return position.y
}

// SetX sets the robot linear velocity to the velocity passed as argument
func (position *Position) SetX(x float64) {
	position.x = x
}

// SetY sets the robot linear velocity to the velocity passed as argument
func (position *Position) SetY(y float64) {
	position.y = y
}
