package world_model

import "sync"

var instance *WorldModel
var once sync.Once

// GetRepositoryInstance gets the singleton instance in the system of commands repository,
// and if it do not exists yet, creates the instance
func GetWorldModelInstance() *WorldModel {
	once.Do(func() {
		instance = newWorldModel()
	})
	return instance
}
