package connection

import (
	"fmt"
	"log"
	"petstore/manage/client"
	"petstore/manage/pet"
	"petstore/manage/tools"

	"github.com/gobuffalo/pop/v6"
)

var db *pop.Connection

func init() {
	pop.Debug = true
	pop.AddLookupPaths("../../")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	var err error
	db, err = pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Conected to database!")
}

func SavePet(Pet pet.Pet) {
	if err := db.Create(&Pet); err != nil {
		log.Fatal("error guardando la nueva mascota", err.Error())
	}

	pets := pet.Pets{}

	if err := db.All(&pets); err != nil {
		log.Fatal("error listando las mascotas", err.Error())
	}

	res, err := tools.PrettyStruct(pets)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Pets", res)
}

func ListAllPets() string {
	pets := pet.Pets{}

	if err := db.All(&pets); err != nil {
		log.Fatal("error consultando todas las mascotas :(", err.Error())
	}

	res, err := tools.PrettyStruct(pets)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func SaveClient(Client client.Client) {
	if err := db.Create(&Client); err != nil {
		log.Fatal("error guardando el client", err.Error())
	}

	clients := client.Clients{}

	if err := db.All(&clients); err != nil {
		log.Fatal("error listando los clientes", err.Error())
	}

	res, err := tools.PrettyStruct(clients)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Clients", res)
}

func ListAllClients() string {
	clients := client.Clients{}

	if err := db.All(&clients); err != nil {
		log.Fatal("error consultando todos los clientes :(", err.Error())
	}

	res, err := tools.PrettyStruct(clients)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
