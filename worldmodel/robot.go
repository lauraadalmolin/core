package world_model

type Robot struct {
	id             int    // robot's id
	role           string // robot's role
	actualPosition Pose   // robot's actual position
	targetPosition Pose   // robot's desired position
}

// SetID sets the robot's id to the value passed as an argument
func (robot *Robot) SetID(id int) {
	robot.id = id
}

// SetRole sets the robot's role to the value passed as an argument
func (robot *Robot) SetRole(role string) {
	robot.role = role
}

// SetActualPosition sets the robot's actualPosition's orientation to
// the float64 value passed as an argument
func (robot *Robot) SetActualPosition(actualPosition Pose) {
	robot.actualPosition = actualPosition
}

// SetTargetPosition sets the robot's targetPosition's orientation to
// the float64 value passed as an argument
func (robot *Robot) SetTargetPosition(targetPosition Pose) {
	robot.targetPosition = targetPosition
}

// GetID returns the robot's id
func (robot Robot) GetID() int {
	return robot.id
}

// GetRole returns the robot's role
func (robot Robot) GetRole() string {
	return robot.role
}

// GetActualPosition returns the robot's actual position
func (robot Robot) GetActualPosition() Pose {
	return robot.actualPosition
}

// GetTargetPosition return the robot's target position
func (robot Robot) GetTargetPosition() Pose {
	return robot.targetPosition
}
