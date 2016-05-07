// Executor project main.go
package main

import (
	"log"
	"os"

	"Executor/fleet"

	"github.com/versus/easyssh"
)

func main() {

	hosts, err := fleet.GetHostsFromYML("Hoth")
	if err != nil {
		log.Fatalln("Fleet Hosts error: %v\n", err)
		os.Exit(1)
	}

	if hosts[0].IsHost() {
		log.Println("Host validate\n")
		log.Println(hosts[0])
	}

	ssh := &easyssh.MakeConfig{
		User:   "versus",
		Server: "hoth.server",
		//Password: "123",
		Key:  "./battle/Hoth/factors/hoth.key",
		Port: "22",
	}

	response, err := ssh.Run("sudo ps ax")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		log.Println(response)
	}

}
