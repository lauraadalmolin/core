package commands

import "sync"

var instance *Repository
var once sync.Once

// GetRepositoryInstance gets the singleton instance in the system of commands repository,
// and if it do not exists yet, creates the instance
func GetRepositoryInstance() *Repository {
	once.Do(func() {
		instance = newRepository()
	})
	return instance
}
