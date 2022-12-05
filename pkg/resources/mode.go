package resources

import (
	"kubedoom/pkg/monsters"
	"kubedoom/pkg/resources/completed"
	"kubedoom/pkg/resources/ingress"
	"kubedoom/pkg/resources/pod"
)

type ClientSet struct {
	Test  *Monster
	Test2 *Entities
}

func (c *ClientSet) Add() monsters.IMonster {
	return &monsters.Monsters{}
}

func (c *ClientSet) Entities() Entities {
	return *c.Entities
}

type Data interface {
	Monster
	Entities
}

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

func (r Resources) Create() Data {
	// monster := new(monsters.Monsters)
	switch r.modeFlag {
	case "ingress":
		return ClientSet{
			Test:  monsters.Monsters{},
			Test2: ingress.Mode{Label: r.labelFlag, Respawn: r.respawnFlag},
		}
	case "completed":
		return ClientSet{
			*Monster:  monsters.Monsters{},
			*Entities: completed.Mode{Label: r.labelFlag, Respawn: r.respawnFlag},
		}
	default:
		return ClientSet{
			*Monster:  monsters.Monsters{},
			*Entities: pod.Mode{Label: r.labelFlag, Respawn: r.respawnFlag},
		}
	}
}
