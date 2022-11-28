package main

import (
	"flag"
	"kubedoom/pkg/doom"
	"kubedoom/pkg/resources"
	"log"
	"net"
)

var (
	modeFlag    string
	labelFlag   string
	respawnFlag bool
)

func init() {
	flag.StringVar(&modeFlag, "mode", "pods", "What to kill pods|ingress")
	flag.StringVar(&labelFlag, "label", "", "labels selected")
	flag.BoolVar(&respawnFlag, "respawn", false, "respawn pod")
	flag.Parse()
}

func main() {
	listener, err := net.Listen("unix", "/dockerdoom.socket")
	if err != nil {
		log.Fatalf("Could not create socket file")
	}
	doom.LaunchVM()
	doom.Run(listener, resources.New(modeFlag, labelFlag, respawnFlag).Create())
}
