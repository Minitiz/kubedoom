package resources

import (
	"kubedoom/pkg/resources/completed"
	"kubedoom/pkg/resources/ingress"
	"kubedoom/pkg/resources/pod"
)

type Type interface {
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

func (r Resources) Create() Type {
	switch r.modeFlag {
	case "ingress":
		return ingress.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}
	case "completed":
		return completed.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}
	default:
		return pod.Mode{Label: r.labelFlag, Respawn: r.respawnFlag}
	}
}
