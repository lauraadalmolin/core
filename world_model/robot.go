package world_model

// import

type Robot struct {
	id             int
	role           string
	actualPosition Pose
	targetPosition Pose
}

func (robot Robot) GetID() float64 {
	return robot.id
}

func (robot Robot) GetRole() float64 {
	return robot.role
}

func (robot Robot) GetActualPosition() Pose {
	return robot.actualPosition.GetPose()
}

func (robot Robot) GetTargetPosition() Pose {
	return robot.targetPosition.GetPose()
}

func (robot Robot) SetID(id int) {
	robot.id = id
}

func (robot Robot) GetRole(role string) {
	position.role = role
}

func (robot Robot) SetActualPosition(actualPosition Pose) Pose {
	position.actualPosition.SetActualPosition(actualPosition)
}

func (robot Robot) SetTargetPosition(targetPosition Pose) Pose {
	position.targetPosition.SetTargetPosition(targetPosition)
}
