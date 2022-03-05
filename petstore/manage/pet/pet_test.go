package pet_test

import (
	"fmt"
	"log"
	"petstore/manage/pet"
	"petstore/manage/tools"
	"testing"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
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
		cat := pet.Pet{
			Animal:    "Cat",
			Price:     10,
			Age:       1,
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

		petsToSave := pet.Pets{
			{
				Animal:    "Dog",
				Price:     10,
				Age:       1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Bunny",
				Price:     8,
				Age:       2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if errs := tx.Create(&petsToSave); errs != nil {
			r.Fail("error guardando nuevas mascotas :(")
		}

		pets := pet.Pets{}

		if err := tx.All(&pets); err != nil {
			r.Fail("error consultando todas las mascotas :(")
		}

		res, err := tools.PrettyStruct(pets)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("All Pets", res)
	})
}

func Test_Find(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		petsToSave := pet.Pets{
			{
				Animal:    "Dog",
				Price:     10,
				Age:       1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Bunny",
				Price:     8,
				Age:       2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Cat",
				Price:     4,
				Age:       2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Cat",
				Price:     4,
				Age:       8,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Cat",
				Price:     4,
				Age:       3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Animal:    "Bird",
				Price:     2,
				Age:       3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		if errs := tx.Create(&petsToSave); errs != nil {
			r.Fail("error guardando una nueva mascota :(")
		}
		pets1 := pet.Pets{}

		// select cats
		q := tx.Where("animal = 'Cat'")

		//age between 1 y 4
		q.Where("age >= (?) AND age <= (?)", 1, 4)

		//execute query
		if err := q.All(&pets1); err != nil {
			r.Fail("error listando los animales que son gatos :(")
		}

		res1, err := tools.PrettyStruct(pets1)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("All Pets First Query Pretty", res1)

		pets2 := pet.Pets{}

		q = tx.Where("price >= (?) AND price <= (?)", 3, 6)

		if err := q.All(&pets2); err != nil {
			r.Fail("error listando los animales que cuestan entre 3 y 6")
		}

		res2, err := tools.PrettyStruct(pets2)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("All Pets Second Query Pretty", res2)
	})
}

func Test_Destroy(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		horse := pet.Pet{
			Animal:    "Horse",
			Price:     25,
			Age:       24,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.Create(&horse); err != nil {
			r.Fail("error guardando una nueva mascota :(")
		}

		if err := tx.Destroy(&horse); err != nil {
			r.Fail("error eliminando la mascota :(")
		}
	})
}

func Test_Update(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		transaction(func(tx *pop.Connection) {
			r := require.New(t)

			tx.RawQuery("INSERT INTO pets (id, animal, price, age, created_at, updated_at) VALUES ('0eb2fd45-fafb-42cc-b182-66a1688b15d7', 'Camel', 40, 4, now(), now())").Exec()

			id := "0eb2fd45-fafb-42cc-b182-66a1688b15d7"

			dog := pet.Pet{
				ID:        uuid.Must(uuid.FromString(id)),
				Animal:    "Dog",
				Price:     35,
				Age:       4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := tx.Update(&dog); err != nil {
				r.Fail("error actualizando la mascotas :(", err.Error())
			}

			res, errr := tools.PrettyStruct(dog)

			if errr != nil {
				log.Fatal(errr.Error())
			}

			fmt.Println("Pet updated", res)

		})
	})
}
