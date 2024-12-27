package main

import (
	"fmt"
	"log"

	"github.com/klambri/ldap-api-bridge/api/routes"
	"github.com/klambri/ldap-api-bridge/internal/configuration"
	ldapconnector "github.com/klambri/ldap-api-bridge/internal/ldap-connector"
)

func main() {

	config, err := configuration.ReadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := ldapconnector.New(config); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ready..")
	routes.Run(config)
}
