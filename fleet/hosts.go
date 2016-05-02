package fleet

import (
	"Executor/r2d2"
	"io/ioutil"
	"log"

	"github.com/versus/yaml"
)

var Hosts []r2d2.Host

func GetHostsFromYML(battlename string) (Hosts []r2d2.Host, err error) {
	file, err := ioutil.ReadFile("./battle/" + battlename + "/factors/hosts.yml")
	if err != nil {
		log.Fatalln("File error: %v\n", err)
		return nil, err
	}
	err = yaml.Unmarshal(file, &Hosts)
	if err != nil {
		log.Fatalf("err: %v\n", err)
		return nil, err
	}

	return Hosts, nil
}
