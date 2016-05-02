package r2d2

import (
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
)

type Host struct {
	Hostname   string `json:"hostname" valid:"alphanum,required"`
	Connection string `json:"connection" valid:"alphanum,required"`
	Address    string `json:"address" valid:"dns,ipv4"`
	Env        string `json:"env" valid:"alphanum`
	User       string `json:"user" valid:"alphanum"`
	Password   string `json:"password" valid:"alphanum`
	Port       string `json:"port" valid:"port" `
	KeyFile    string `json:"key" valid:"alphanum`
	Desc       string `json:"desc" valid:"alphanum`
}

func (h *Host) IsHost() bool {
	res, err := govalidator.ValidateStruct(h)
	if err != nil {
		log.Fatalf("error: %v\n", err.Error())
	}
	return res
}

func (h *Host) String() string {
	return fmt.Sprintf("\n Hostname: %v \n Connection: %v \n", h.Hostname, h.Connection)

}
