package commands

// Command - command to send
type Command struct {
	robotID         int
	linearVelocity  float64
	angularVelocity float64
}

// GetRobotID returns the robot identification
func (command Command) GetRobotID() int {
	return command.robotID
}

// GetLinearVelocity returns the robot linear velocity
func (command Command) GetLinearVelocity() float64 {
	return command.linearVelocity
}

// GetAngularVelocity returns the robot angular velocity
func (command Command) GetAngularVelocity() float64 {
	return command.angularVelocity
}

// SetRobotID sets the robot identification to the ID passed as argument
func (command *Command) SetRobotID(id int) {
	command.robotID = id
}

// SetLinearVelocity sets the robot linear velocity to the velocity passed as argument
func (command *Command) SetLinearVelocity(velocity float64) {
	command.linearVelocity = velocity
}

// SetAngularVelocity sets the robot angular velocity to the velocity passed as argument
func (command *Command) SetAngularVelocity(velocity float64) {
	command.angularVelocity = velocity
}
