package world

type Entity interface {
	GetID() string
}

type World struct {
	ID       string
	Entities []Entity
}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}

func (w *World) ChangeEntityDimension(e Entity, world *World) {
	for i, entity := range w.Entities {
		if entity.GetID() == e.GetID() {
			w.Entities = append(w.Entities[:i], w.Entities[i+1:]...)
			break
		}
	}

	world.AddEntity(e)
}
