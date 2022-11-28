package ingress

import "strings"

func (m Mode) Format(entity string) string {
	return strings.Split(entity, "/")[1]
}
