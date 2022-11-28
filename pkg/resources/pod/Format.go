package pod

import "strings"

func (m Mode) Format(entity string) string {
	podparts := strings.Split(entity, "/")[1]
	splited := strings.Split(podparts, "-")
	removed := splited[:len(splited)-2]
	return strings.Join(removed, "-")
}
