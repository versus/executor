// Executor project main.go
package main

import (
	"log"
	"os"

	"Executor/fleet"
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

}
