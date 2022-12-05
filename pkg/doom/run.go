package doom

import (
	"kubedoom/pkg/resources"
	"log"
	"net"
	"strconv"
	"strings"
)

func hash(input string) int32 {
	var hash int32
	hash = 5381
	for _, char := range input {
		hash = (hash << 5) + hash + int32(char)
	}
	if hash < 0 {
		hash = 0 - hash
	}
	return hash
}

func Run(listener net.Listener, mode resources.Data) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		stop := false
		for !stop {
			log.Print(mode)
			bytes := make([]byte, 40960)
			n, err := conn.Read(bytes)
			if err != nil {
				stop = true
			}
			bytes = bytes[0:n]
			strbytes := strings.TrimSpace(string(bytes))
			entities := mode.GetEntities(MONSTERS)
			if strbytes == "list" {
				for _, entity := range entities {
					log.Print(entity)
					padding := strings.Repeat("\n", 255-len(entity))
					_, err = conn.Write([]byte(mode.Format(entity) + padding))
					if err != nil {
						log.Fatal("Could not write to socker file")
					}
				}
				MONSTERS = entities
				conn.Close()
				stop = true
			} else if strings.HasPrefix(strbytes, "kill ") {
				parts := strings.Split(strbytes, " ")
				killhash, err := strconv.ParseInt(parts[1], 10, 32)
				if err != nil {
					log.Fatal("Could not parse kill hash")
				}
				for index, entity := range entities {
					if hash(entity) == int32(killhash) {
						mode.DeleteEntity(entity)
						MONSTERS = append(entities[:index], entities[index+1:]...)
						break
					}
				}
				conn.Close()
				stop = true
			}
		}
	}
}
