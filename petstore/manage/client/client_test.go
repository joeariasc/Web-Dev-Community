package client_test

import (
	"fmt"
	"log"
	"petstore/manage/client"
	"petstore/manage/tools"
	"testing"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
)

var db *pop.Connection

func init() {
	pop.Debug = true
	pop.AddLookupPaths("../../")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	var err error
	db, err = pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}
}

func transaction(fn func(tx *pop.Connection)) {
	err := db.Rollback(func(tx *pop.Connection) {
		fn(tx)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Create(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		cat := client.Client{
			Name:      "Joe",
			LastName:  "Arias",
			Email:     "joe@petstore.com",
			Phone:     "333000333",
			Address:   "fake street",
			Gender:    "Male",
			Age:       20,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&cat); err != nil {
			r.Fail("error guardando una nueva mascota :(", err.Error())
		}
	})
}

func Test_List(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)

		clientsToSave := client.Clients{
			{
				Name:      "Joe",
				LastName:  "Arias",
				Email:     "joe@petstore.com",
				Phone:     "333000333",
				Address:   "fake street",
				Gender:    "Male",
				Age:       20,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Name:      "Luis",
				LastName:  "Pedraza",
				Email:     "luis@petstore.com",
				Phone:     "3330002222",
				Address:   "fake house",
				Gender:    "Male",
				Age:       22,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if errs := tx.Create(&clientsToSave); errs != nil {
			r.Fail("error guardando los clientes :(")
		}

		clients := client.Clients{}

		if err := tx.All(&clients); err != nil {
			r.Fail("error consultando todos los clientes :(")
		}

		res, err := tools.PrettyStruct(clients)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("All Clients", res)
	})
}
