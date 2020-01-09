package commands

// Repository - type to store the commands to send
type Repository [3]Command

func newRepository() *Repository {
	return &Repository{}
}
