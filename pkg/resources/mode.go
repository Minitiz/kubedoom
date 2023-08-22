package resources

import (
	"kubedoom/pkg/monsters"
	"kubedoom/pkg/resources/completed"
	"kubedoom/pkg/resources/ingress"
	"kubedoom/pkg/resources/pod"
)

type ClientSet struct {
	Monster  *Monster
	Entities *Entities
}

type Data map[string]Entities

type Monster interface {
	Add(monsters.Monster)
	Delete(monsters.Monster)
}

type Entities interface {
	GetEntities([]string) []string
	DeleteEntity(string)
	Format(string) string
}

type Resources struct {
	modeFlag    string
	labelFlag   string
	respawnFlag bool
}

func New(modeFlag string, labelFlag string, respawnFlag bool) Resources {
	return Resources{
		modeFlag:    modeFlag,
		labelFlag:   labelFlag,
		respawnFlag: respawnFlag,
	}
}

func (r Resources) Create() Entities {
	// monster := new(monsters.Monsters)

	var data Data = make(Data)
	data["ingress"] = &ingress.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}
	data["completed"] = &completed.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}
	data["pod"] = &pod.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}

	return data[r.modeFlag]
}
