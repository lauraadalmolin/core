package world_model

type WorldModel struct {
	Ball Ball
	Team Team
}

func newWorldModel() *WorldModel {
	return &WorldModel{}
}
