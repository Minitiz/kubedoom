package ingress

import (
	"log"
	"os/exec"
	"strings"
)

func (m Mode) DeleteEntity(entity string) {
	log.Printf("Pod to kill: %v", entity)
	podparts := strings.Split(entity, "/")
	cmd := exec.Command("/usr/bin/kubectl", "delete", "ingress", "-n", podparts[0], podparts[1])
	go cmd.Run()
}
