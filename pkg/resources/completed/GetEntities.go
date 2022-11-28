package completed

import (
	"log"
	"os"
	"os/exec"
	"regexp"
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

var toddBatch = regexp.MustCompile(`todd-[a-z1-9-]*-batch`)
var status = regexp.MustCompile(`(Failed|Succeeded)`)

func sortJob(all []string) []string {
	ret := []string{}
	for _, entry := range all {
		if toddBatch.MatchString(entry) && status.MatchString(entry) {
			ret = append(ret, strings.Join(strings.Split(entry, "/")[:2], "/"))
		}
	}
	return ret
}

func (m Mode) GetEntities(MONSTERS []string) []string {
	if m.Respawn && firstIteration {
		return MONSTERS
	}
	args := []string{"kubectl", "get", "pods", "-l", m.Label, "-A", "-o", "go-template", "--template={{range .items}}{{.metadata.namespace}}/{{.metadata.name}}/{{.status.phase}} {{end}}"}
	output := outputCmd(args)
	outputstr := strings.TrimSpace(output)
	firstIteration = true
	log.Print(sortJob(strings.Split(outputstr, " ")))
	return sortJob(strings.Split(outputstr, " "))
}
