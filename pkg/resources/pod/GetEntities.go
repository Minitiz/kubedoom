package pod

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func outputCmd(argv []string) string {
	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("The following command failed: \"%v\"\n", argv)
	}
	return string(output)
}

func (m Mode) GetEntities(MONSTERS []string) []string {
	if m.Respawn && firstIteration {
		return MONSTERS
	}
	var args []string
	if namespace, exists := os.LookupEnv("NAMESPACE"); exists {
		args = []string{"kubectl", "get", "pods", "-l", m.Label, "--namespace", namespace, "-o", "go-template", "--template={{range .items}}{{.metadata.namespace}}/{{.metadata.name}} {{end}}"}
	} else {
		args = []string{"kubectl", "get", "pods", "-A", "-o", "go-template", "--template={{range .items}}{{.metadata.namespace}}/{{.metadata.name}} {{end}}"}
	}
	output := outputCmd(args)
	outputstr := strings.TrimSpace(output)
	firstIteration = true
	return strings.Split(outputstr, " ")
}
