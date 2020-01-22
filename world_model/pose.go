package world_model

type Pose struct {
	orientation float64
}

// GetOrientation gets the orientation of a Pose
func (pose Pose) GetOrientation() float64 {
	return pose.orientation
}

// SetOrientation sets the orientation of a Pose
func (pose *Pose) SetOrientation(orientation float64) {
	pose.orientation = orientation
}
